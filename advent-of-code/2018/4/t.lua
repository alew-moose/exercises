#!/usr/bin/env luajit

-- _G.p = require("pimp")

function get_input(file)
    local events = {}
    for line in io.lines(file) do
        event = parse_event(line) or error(string.format("failed to parse %q", line))
        table.insert(events, event)
    end
    table.sort(events, function(e1, e2) return e1.ts < e2.ts end)
    return events
end

function parse_event(s)
    local ts_pattern = "^%[(%d%d%d%d%-%d%d%-%d%d) (%d%d):(%d%d)%] (.+)$"
    local date, hh, mm, rest = string.match(s, ts_pattern)
    if date == nil then
        return nil
    end
    local guard_id, event_type
    guard_id = string.match(rest, "^Guard #(%d+) begins shift")
    if guard_id ~= nil then
        event_type = "begin_shift"
    elseif rest == "wakes up" then
        event_type = "wake_up"
    elseif rest == "falls asleep" then
        event_type = "fall_asleep"
    else
        return nil
    end
    return {
        event_type = event_type,
        ts = date .. " " .. hh .. ":" .. mm,
        date = date,
        h = tonumber(hh),
        m = tonumber(mm),
        guard_id = guard_id,
    }
end

function make_sleep_intervals(events)
    local intervals = {}
    local guard_id, asleep, date, date_intervals, interval
    for _, e in ipairs(events) do
        if e.event_type == "begin_shift" then
            if asleep then
                error(string.format("[%s] guard %d is asleep", e.ts, guard_id))
            end
            guard_id = e.guard_id
            asleep = false
        elseif e.event_type == "fall_asleep" then
            if asleep then
                error(string.format("[%s] guard %d already asleep", e.ts, guard_id))
            end
            asleep = true
            interval = {
                guard_id = guard_id,
                date = e.date,
                start_m = e.m
            }
        elseif e.event_type == "wake_up" then
            if not asleep then
                error(string.format("[%s] guard %d is not asleep", e.ts, guard_id))
            end
            if e.date ~= interval.date then
                error(string.format("[%s] dates don't match: %q %q", e.ts, e.date, interval.date))
            end
            asleep = false
            interval.end_m = e.m - 1
            table.insert(intervals, interval)
        else
            error(string.format("invalid event_type %q", e.event_type or "<nil>"))
        end
    end
    return intervals
end

function merge_intervals_by_date(intervals)
    local merged = {}
    local date_intervals
    for _, interval in ipairs(intervals) do
        if date_intervals == nil or date_intervals.date ~= interval.date then
            if date_intervals ~= nil then
                table.insert(merged, date_intervals)
                date_intervals = nil
            end
            date_intervals = {
                date = interval.date,
                guard_id = interval.guard_id,
                intervals = {
                    {
                        start_m = interval.start_m,
                        end_m = interval.end_m,
                    },
                },
            }
        else
            table.insert(date_intervals.intervals, {
                start_m = interval.start_m,
                end_m = interval.end_m,
            })
        end
    end
    if date_intervals then
        table.insert(merged, date_intervals)
    end
    return merged
end

function merge_intervals_by_guard(date_intervals)
    local merged = {}
    for _, di in pairs(date_intervals) do
        if not merged[di.guard_id] then
            merged[di.guard_id] = {}
        end
        for _, interval in ipairs(di.intervals) do
            table.insert(merged[di.guard_id], interval)
        end
    end
    return merged
end

function solve_part_1(events)
    local sleep_intervals = make_sleep_intervals(events)
    local date_intervals = merge_intervals_by_date(sleep_intervals)
    local guard_intervals = merge_intervals_by_guard(date_intervals)

    local guard_sleep_total = {}
    for _, di in pairs(date_intervals) do
        local sleep_sum = 0
        for _, interval in ipairs(di.intervals) do
            sleep_sum = sleep_sum + (interval.end_m - interval.start_m + 1)
        end
        guard_sleep_total[di.guard_id] = (guard_sleep_total[di.guard_id] or 0) + sleep_sum
    end

    local max_guard_id = -1
    local max_sleep_total = -1
    for guard_id, sleep_total in pairs(guard_sleep_total) do
        if sleep_total > max_sleep_total then
            max_sleep_total = sleep_total
            max_guard_id = guard_id
        end
    end
    if max_guard_id == -1 then
        error("no max_guard_id")
    end

    local minutes = {}
    for _, interval in ipairs(guard_intervals[max_guard_id]) do
        for m = interval.start_m, interval.end_m do
            minutes[m] = (minutes[m] or 0) + 1
        end
    end

    local max_minute = -1
    local max_sleep = -1
    for minute = 0, 59 do
        if (minutes[minute] or -1) > max_sleep then
            max_minute = minute
            max_sleep = minutes[minute]
        end
    end

    return max_guard_id * max_minute
end

function solve_part_2(events)
    local sleep_intervals = make_sleep_intervals(events)
    local date_intervals = merge_intervals_by_date(sleep_intervals)
    local guard_intervals = merge_intervals_by_guard(date_intervals)

    local guard_minutes = {}
    for guard_id, intervals in pairs(guard_intervals) do
        guard_minutes[guard_id] = {}
        for _, interval in ipairs(intervals) do
            for m = interval.start_m, interval.end_m do
                guard_minutes[guard_id][m] = (guard_minutes[guard_id][m] or 0) + 1
            end
        end
    end

    local max_guard_id = -1
    local max_sleep = -1
    local max_minute = -1
    for guard_id, minutes in pairs(guard_minutes) do
        for minute, sleep in pairs(minutes) do
            if sleep > max_sleep then
                max_sleep = sleep
                max_minute = minute
                max_guard_id = guard_id
            end
        end
    end

    return max_guard_id * max_minute
end


-- local events = get_input("test-input.txt")
local events = get_input("input.txt")

print("part 1:", solve_part_1(events))
print("part 2:", solve_part_2(events))
