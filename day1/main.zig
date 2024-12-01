const std = @import("std");

fn part1(leftSlice: []i64, rightSlice: []i64) u64 {
    var sum : u64 = 0;
    for (leftSlice, rightSlice) |left, right| {
        sum += @abs(left - right);
    }
    return sum;
}

fn part2(leftSlice: []i64, rightSlice: []i64) u64 {
    var sum : u64 = 0;
    var rightIndex : usize = 0;
    for (leftSlice) |left| {
        while (rightIndex < rightSlice.len and rightSlice[rightIndex] < left) {
            rightIndex += 1;
        }
        while (rightIndex < rightSlice.len and rightSlice[rightIndex] == left) {
            sum += @intCast(left);
            rightIndex += 1;
        }
    }
    return sum;
}

pub fn main() !void {
    var alloc = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = alloc.allocator();
    const file = try std.fs.cwd().openFile("input.txt", .{});
    const file_size = (try file.stat()).size;
    const input = try file.readToEndAlloc(allocator, file_size);
    defer allocator.free(input);
    file.close();
    var lines = std.mem.tokenizeAny(u8, input, "\n");
    var leftInts = std.ArrayList(i64).init(allocator);
    var rightInts = std.ArrayList(i64).init(allocator);
    while (lines.next()) |line| {
        var leftAndRight = std.mem.tokenizeScalar(u8, line, ' ');
        const left = leftAndRight.next().?;
        const right = leftAndRight.next().?;
        try leftInts.append(try std.fmt.parseInt(i64, left, 10));
        try rightInts.append(try std.fmt.parseInt(i64, right, 10));
    }
    const leftSlice = try leftInts.toOwnedSlice();
    const rightSlice = try rightInts.toOwnedSlice();
    defer allocator.free(leftSlice);
    defer allocator.free(rightSlice);
    std.mem.sort(i64, leftSlice, {}, comptime std.sort.asc(i64));
    std.mem.sort(i64, rightSlice, {}, comptime std.sort.asc(i64));
    std.debug.print("{}\n", .{part1(leftSlice, rightSlice)});
    std.debug.print("{}\n", .{part2(leftSlice, rightSlice)});
}
