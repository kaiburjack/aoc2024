const std = @import("std");

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

fn part1(field: [][]const u8) void {
    var count : u64 = 0;
    for (0..field.len) |row| {
        for (0..field[row].len) |col| {
            if (field[row][col] != 'X') {
                continue;
            }
            var dr : i8 = -1;
            while (dr <= 1) : (dr += 1) {
                var dc : i8 = -1;
                while (dc <= 1) : (dc += 1) {
                    if (dr == 0 and dc == 0) {
                        continue;
                    }
                    if (find("XMAS", field, @intCast(row), @intCast(col), dr, dc)) {
                        count += 1;
                    }
                }
            }
        }
    }
    std.debug.print("{}\n", .{count});
}

fn part2(field: [][]const u8) void {
    var count: u64 = 0;
    var row : i32 = 0;
    while (row < field.len) : (row += 1) {
        var col : i32 = 0;
        while (col < field[@intCast(row)].len) : (col += 1) {
            if (field[@intCast(row)][@intCast(col)] != 'A') {
                continue;
            }
            if (find("MAS", field, row-1, col-1, 1, 1) and
                find("MAS", field, row+1, col-1, -1, 1) or
                find("MAS", field, row+1, col-1, -1, 1) and
                    find("MAS", field, row+1, col+1, -1, -1) or
                find("MAS", field, row-1, col-1, 1, 1) and
                    find("MAS", field, row-1, col+1, 1,-1) or
                find("MAS", field, row-1, col+1, 1, -1) and
                    find("MAS", field, row+1, col+1, -1, -1)) {
                count += 1;
            }
        }
    }
    std.debug.print("{}\n", .{count});
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

    part1(lines.items);
    part2(lines.items);
}
