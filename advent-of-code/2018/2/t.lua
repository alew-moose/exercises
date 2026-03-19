#!/usr/bin/env luajit

local function get_input(file)
    local lines = {}
    for line in io.lines(file) do
        table.insert(lines, line)
    end
    return lines
end


local function count_bytes(s)
    local bytes = {}
    for i = 1, string.len(s) do
        local b = string.byte(s, i)
        bytes[b] = (bytes[b] or 0) + 1
    end
    return bytes
end

local function repeated(s)
    local r = {[2] = false, [3] = false}
    local bytes = count_bytes(s)
    for _, cnt in pairs(bytes) do
        if cnt == 2 then
            r[2] = true
        elseif cnt == 3 then
            r[3] = true
        end
    end
    return r
end

local function solve_part_1(lines)
    local r2, r3 = 0, 0
    for i, line in ipairs(lines) do
        local r = repeated(line)
        if r[2] then
            r2 = r2 + 1
        end
        if r[3] then
            r3 = r3 + 1
        end
    end
    return r2 * r3
end

local function solve_part_2(lines)
    for i1 = 1, #lines-1 do
        local l1 = lines[i1]
        for i2 = i1+1, #lines do
            local l2 = lines[i2]
            local diff_cnt = 0
            local diff_i = -1
            for i = 1, string.len(l1) do
                if string.sub(l1, i, i) ~= string.sub(l2, i, i) then
                    diff_cnt = diff_cnt + 1
                    diff_i = i
                end
                if diff_cnt > 1 then
                    goto next_line2
                end
            end
            if diff_cnt == 1 then
                return string.sub(l1, 1, diff_i-1) .. string.sub(l1, diff_i+1, -1)
            end
            ::next_line2::
        end
    end
    return "-"
end


local lines = get_input("input.txt")

print("part 1:", solve_part_1(lines))
print("part 2:", solve_part_2(lines))
