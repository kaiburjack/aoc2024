const std = @import("std");

const Direction = enum {
    Undefined,
    Increase,
    Decrease,
};

const XMAS = "XMAS";
const MAS = "MAS";

fn find(pattern: []const u8, field: [][]const u8, row : i32, col : i32, dr : i8, dc: i8) bool {
    var j : usize = 0;
    var c : i32 = col;
    var r : i32 = row;
    while (j < pattern.len and r >= 0 and r < field.len and c >= 0 and c < field[@intCast(row)].len) : ({c +%= dc; r +%= dr;}) {
        if (field[@intCast(r)][@intCast(c)] != pattern[j]) {
            return false;
        }
        j += 1;
    }
    return j == pattern.len;
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();
    const file = try std.fs.cwd().openFile("real.txt", .{});
    const file_size = (try file.stat()).size;
    const input = try file.readToEndAllocOptions(allocator, file_size, file_size, 8, null);
    file.close();
    var lines_iterator = std.mem.tokenizeScalar(u8, input, '\n');
    var lines = std.ArrayList([]const u8).init(allocator);
    while (lines_iterator.next()) |line| {
        try lines.append(line);
    }
    var count : u64 = 0;
    // part 1
    for (0..lines.items.len) |row| {
        for (0..lines.items[row].len) |col| {
            if (lines.items[row][col] != 'X') {
                continue;
            }
            var dr : i8 = -1;
            while (dr <= 1) : (dr += 1) {
                var dc : i8 = -1;
                while (dc <= 1) : (dc += 1) {
                    if (dr == 0 and dc == 0) {
                        continue;
                    }
                    if (find(XMAS, lines.items, @intCast(row), @intCast(col), dr, dc)) {
                        count += 1;
                    }
                }
            }
        }
    }
    std.debug.print("Count: {}\n", .{count});

    // part 2
    count = 0;
    var row : i32 = 0;
    while (row < lines.items.len) : (row += 1) {
        var col : i32 = 0;
        while (col < lines.items[@intCast(row)].len) : (col += 1) {
            if (lines.items[@intCast(row)][@intCast(col)] != 'A') {
                continue;
            }
            if (find(MAS, lines.items, row-1, col-1, 1, 1) and
                find(MAS, lines.items, row+1, col-1, -1, 1) or
                find(MAS, lines.items, row+1, col-1, -1, 1) and
                find(MAS, lines.items, row+1, col+1, -1, -1) or
                find(MAS, lines.items, row-1, col-1, 1, 1) and
                find(MAS, lines.items, row-1, col+1, 1,-1) or
                find(MAS, lines.items, row-1, col+1, 1, -1) and
                find(MAS, lines.items, row+1, col+1, -1, -1)) {
                count += 1;
            }
        }
    }
    std.debug.print("Count: {}\n", .{count});
}
