#!/usr/bin/env luajit

function get_input(file)
    local rects = {}
    for line in io.lines(file) do
        local rect = parse_rect(line)
        if rect then
            table.insert(rects, rect)
        else
            error(string.format("failed to parse %q", line))
        end
    end
    return rects
end

function parse_rect(s)
    local pattern = "^#(%d+) @ (%d+),(%d+): (%d+)x(%d+)$"
    local id, x_str, y_str, w_str, h_str = string.match(s, pattern)
    if id == nil then
        return nil
    end
    return {
        id = id,
        x = x_str + 1,
        y = y_str + 1,
        w = tonumber(w_str), -- width
        h = tonumber(h_str), -- height
    }
end

function get_dimensions(rects)
    local w, h = 0, 0
    for _, r in ipairs(rects) do
        local right = r.x + r.w - 1
        if right > w then
            w = right
        end
        local bottom = r.y + r.h - 1
        if bottom > h then
            h = bottom
        end
    end
    return w, h
end

function make_map(w, h)
    local map = {}
    for y = 1, h do
        local row = {}
        for x = 1, w do
            table.insert(row, 0)
        end
        table.insert(map, row)
    end
    return map
end

function print_map(map)
    for y = 1, #map do
        local row = map[y]
        for x = 1, #row do
            local c = '.'
            if map[y][x] == 1 then
                c = 'o'
            elseif map[y][x] > 1 then
                c = 'x'
            end
            io.stdout:write(c)
        end
        io.stdout:write("\n")
    end
end

function add_to_map(map, rect)
    for y = rect.y, rect.y+rect.h-1 do
        for x = rect.x, rect.x+rect.w-1 do
            map[y][x] = (map[y][x] or 0) + 1
        end
    end
end

function solve_part_1(rects)
    local w, h = get_dimensions(rects)
    local map = make_map(w, h)
    for _, rect in ipairs(rects) do
        add_to_map(map, rect)
    end
    local sum = 0
    for y = 1, #map do
        for x = 1, #map[1] do
            if map[y][x] > 1 then
                sum = sum + 1
            end
        end
    end
    return sum
end

function solve_part_2(rects)
    local w, h = get_dimensions(rects)
    local map = make_map(w, h)
    for _, rect in ipairs(rects) do
        add_to_map(map, rect)
    end
    for _, r in ipairs(rects) do
        for y = r.y, r.y+r.h-1 do
            for x = r.x, r.x+r.w-1 do
                if map[y][x] > 1 then
                    goto next_rect
                end
            end
        end
        do return r.id end
        ::next_rect::
    end
    return -1
end




-- local rects = get_input("test-input.txt")
local rects = get_input("input.txt")

print("part 1:", solve_part_1(rects))
print("part 2:", solve_part_2(rects))

