const std = @import("std");
const stdout = std.io.getStdOut().writer();

pub fn main() !void {
    const gpa = std.heap.page_allocator;

    var biglist = std.ArrayList(u8).init(gpa);
    defer biglist.deinit();

    // for (biglist[0..biglist.len]) |item| {
    //
    // }



    var index: u32 = 1;
    while(index < 100000000) :(index += 1) {
        if (index % 10 == 7  or index % 7 == 0){
            try biglist.writer().print("SMAC\n", .{});
        }else{
            try biglist.writer().print("{d}\n", .{index});
        }
    }

    try print_list(&biglist);
}

fn print_list(list: *std.ArrayList(u8)) !void {
    var iter = std.mem.splitSequence(u8, list.items, "\n");
    var count: usize = 0;
    while (iter.next()) |item| {
        if (item.len > 0) {
            count += 1;
        }
    }
    std.debug.print("Found {d} items\n", .{count});
}
