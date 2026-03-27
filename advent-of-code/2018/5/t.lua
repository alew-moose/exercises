#!/usr/bin/env luajit

function get_input(file)
    local input
    for line in io.lines(file) do
        input = line
    end
    return input
end

function solve_part_1(s)
    local list = string_to_list(s)
    list = list_remove_pairs(list)
    return list_length(list)
end

function solve_part_2(s)
    local min_len = #s
    for n = string.byte('A'), string.byte('Z') do
        local uc = string.char(n)
        local lc = string.char(n + 32)
        local pattern = string.format("[%s%s]", uc, lc)
        local copy = s
        copy = string.gsub(copy, pattern, "")
        local list = string_to_list(copy)
        list = list_remove_pairs(list)
        local s2 = list_to_string(list)
        if #s2 < min_len then
            min_len = #s2
        end
    end
    return min_len
end

function list_remove_pairs(list)
    local head = list
    while head and head.next do
        if math.abs(head.val - head.next.val) == 32 then
            if head.prev then
                head.prev.next = head.next.next
                if head.next.next then
                    head.next.next.prev = head.prev
                end
                head = head.prev
            else
                if head.next.next then
                    head.next.next.prev = nil
                end
                head = head.next.next
                list = head
            end
        else
            head = head.next
        end
    end
    return list
end

function string_to_list(s)
    local list
    local head
    for i = 1, #s do
        local c = s:sub(i, i)
        local node = { val = string.byte(c) }
        if head == nil then
            head = node
            list = head
        else
            head.next = node
            node.prev = head
            head = node
        end
    end
    return list
end

function list_to_string(list)
    local vals = {}
    while list do
        table.insert(vals, string.char(list.val))
        list = list.next
    end
    return table.concat(vals)
end

function list_length(list)
    local len = 0
    while list do
        len = len + 1
        list = list.next
    end
    return len
end


-- input = get_input("test-input.txt")
input = get_input("input.txt")
print("part 1:", solve_part_1(input))
print("part 2:", solve_part_2(input))





