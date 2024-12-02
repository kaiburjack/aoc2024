const std = @import("std");

const Direction = enum {
    Undefined,
    Increase,
    Decrease,
};

fn part1(numbers : []const i64, ignore : ?usize) bool {
    var increaseOrDecrease : Direction = .Undefined;
    var last : ?i64 = null;
    for (0..numbers.len) |i| {
        if (i == ignore) {
            continue;
        } else if (last == null) {
            last = numbers[i];
            continue;
        }
        const number = numbers[i];
        if (increaseOrDecrease == .Undefined) {
            if (number > last.?) {
                increaseOrDecrease = .Increase;
            } else if (number < last.?) {
                increaseOrDecrease = .Decrease;
            } else {
                return false;
            }
        } else if (increaseOrDecrease == .Increase and number <= last.?) {
            return false;
        } else if (increaseOrDecrease == .Decrease and number >= last.?) {
            return false;
        }
        if (@abs(last.?-number) == 0 or @abs(last.?-number) > 3) {
            return false;
        }
        last = number;
    }
    return true;
}

fn part2(numbers : []const i64) bool {
    for (0..numbers.len) |i| {
        if (part1(numbers, i)) {
            return true;
        }
    }
    return false;
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();
    const file = try std.fs.cwd().openFile("real.txt", .{});
    const file_size = (try file.stat()).size;
    const input = try file.readToEndAllocOptions(allocator, file_size, file_size, 8, null);
    file.close();
    var lines = std.mem.tokenizeScalar(u8, input, '\n');
    var numbers = std.ArrayList(i64).init(allocator);
    var safe_lines: i32 = 0;
    var safe_lines_with_dampener: i32 = 0;
    while (lines.next()) |line| {
        var splitted = std.mem.tokenizeScalar(u8, line, ' ');
        try numbers.resize(0);
        while (splitted.next()) |number| {
            const parser = try std.fmt.parseInt(i64, number, 10);
            try numbers.append(parser);
        }
        if (part1(numbers.items, null)) {
            safe_lines += 1;
        }
        if (part2(numbers.items)) {
            safe_lines_with_dampener += 1;
        }
    }
    std.debug.print("{}\n", .{safe_lines});
    std.debug.print("{}\n", .{safe_lines_with_dampener});
}
