const std = @import("std");

const Direction = enum {
    Undefined,
    Increase,
    Decrease,
};

const MAX_DIFFERENCE = 3;

fn is_safe(numbers : []const i64, ignore : ?usize) bool {
    var increaseOrDecrease : Direction = .Undefined;
    var last : i64 = 0;
    var has_last = false;
    for (0..numbers.len) |i| {
        if (i == ignore) {
            continue;
        } else if (!has_last) {
            last = numbers[i];
            has_last = true;
            continue;
        }
        const number = numbers[i];
        switch (increaseOrDecrease) {
            .Undefined => {
                if (number > last and @abs(last-number) <= MAX_DIFFERENCE) {
                    increaseOrDecrease = .Increase;
                } else if (number < last and @abs(last-number) <= MAX_DIFFERENCE) {
                    increaseOrDecrease = .Decrease;
                } else {
                    return false;
                }
            },
            .Increase => if (number <= last or @abs(last-number) > MAX_DIFFERENCE) {
                return false;
            },
            .Decrease => if (number >= last or @abs(last-number) > MAX_DIFFERENCE) {
                return false;
            },
        }
        last = number;
    }
    return true;
}

fn is_safe_dampened(numbers : []const i64) bool {
    for (0..numbers.len) |i| {
        if (is_safe(numbers, i)) {
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
            try numbers.append(try std.fmt.parseInt(i64, number, 10));
        }
        safe_lines += if (is_safe(numbers.items, null)) 1 else 0;
        safe_lines_with_dampener += if (is_safe_dampened(numbers.items)) 1 else 0;
    }
    std.debug.print("{}\n", .{safe_lines});
    std.debug.print("{}\n", .{safe_lines_with_dampener});
}
