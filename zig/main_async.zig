const std = @import("std");
const stdout = std.io.getStdOut().writer();

pub fn main() !void {
    const gpa = std.heap.page_allocator;

    const threadCount = 4;
    var t: [threadCount]std.Thread = undefined;

    var list : [threadCount]std.ArrayList(u8) = undefined;

    var i : u16 = 0;
    while (i < threadCount) : (i += 1){
        list[i] = std.ArrayList(u8).init(gpa);
        defer list[i].clearAndFree();
        // check for threadCount alternative
        t[i] = try std.Thread.spawn(.{},  thread_main, .{i, &list[i]});
    }

    i = 0;
    while(i < threadCount) : (i += 1){
        t[i].join();
    }
}

fn thread_main(id: u32, list: *std.ArrayList(u8)) !void{
    const stepper = 25000000;
    var addr: usize = 0;
    if (id == 3) addr += 1; // 3 thread end at 9 and one at 0

    var index: usize = id * stepper;
    while(index < (id * stepper + stepper + addr)) : (index += 1){
        if (index % 10 == 7  or index % 7 == 0){
            try list.writer().print("SMAC\n", .{});
        }else{
            try list.writer().print("{d}\n", .{index});
        }
    }
    try print_list(list);
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

