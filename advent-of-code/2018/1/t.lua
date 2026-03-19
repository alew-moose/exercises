#!/usr/bin/env luajit

local function get_input(file)
    local nums = {}
    for line in io.lines(file) do
        num = tonumber(line) or error(string.format("failed to parse %q", line))
        table.insert(nums, num)
    end
    return nums
end

local function solve_part_1(nums)
    local sum = 0
    for _, num in ipairs(nums) do
        sum = sum + num
    end
    return sum
end

local function solve_part_2(nums)
    local sums = {}
    local sum = 0
    sums[sum] = true
    while true do
        for _, num in ipairs(nums) do
            sum = sum + num
            if sums[sum] then
                return sum
            end
            sums[sum] = true
        end
    end
end


local nums = get_input("input.txt")

print("part 1:", solve_part_1(nums))
print("part 2:", solve_part_2(nums))
