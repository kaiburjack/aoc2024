const std = @import("std");

fn part1(left_slice: []const i64, right_slice: []const i64) u64 {
    var sum : u64 = 0;
    for (left_slice, right_slice) |left, right| {
        sum += @abs(left - right);
    }
    return sum;
}

fn part2(left_slice: []const i64, right_slice: []const i64) u64 {
    var sum : u64 = 0;
    var right_index : usize = 0;
    for (left_slice) |left| {
        while (right_index < right_slice.len and right_slice[right_index] < left) : (right_index += 1) {
        }
        while (right_index < right_slice.len and right_slice[right_index] == left) : (right_index += 1) {
            sum += @intCast(left);
        }
    }
    return sum;
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();
    const file = try std.fs.cwd().openFile("input.txt", .{});
    const file_size = (try file.stat()).size;
    const input = try file.readToEndAllocOptions(allocator, file_size, file_size, 8, null);
    file.close();
    var lines = std.mem.tokenizeScalar(u8, input, '\n');
    var left_ints = std.ArrayList(i64).init(allocator);
    var right_ints = std.ArrayList(i64).init(allocator);
    while (lines.next()) |line| {
        var left_and_right = std.mem.tokenizeScalar(u8, line, ' ');
        const left = left_and_right.next().?;
        const right = left_and_right.next().?;
        try left_ints.append(try std.fmt.parseInt(i64, left, 10));
        try right_ints.append(try std.fmt.parseInt(i64, right, 10));
    }
    const left_slice = try left_ints.toOwnedSlice();
    const right_slice = try right_ints.toOwnedSlice();
    std.mem.sort(i64, left_slice, {}, std.sort.asc(i64));
    std.mem.sort(i64, right_slice, {}, std.sort.asc(i64));
    std.debug.print("{}\n", .{part1(left_slice, right_slice)});
    std.debug.print("{}\n", .{part2(left_slice, right_slice)});
}
