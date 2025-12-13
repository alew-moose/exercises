#!/usr/bin/env perl
use v5.42;
use utf8;
use DDP;
use List::Util qw(max);

my $reg_re = qr/[a-z]+/;
my $num_re = qr/-?\d+/;
my $op_re = qr/inc|dec/;
my $cmp_re = qr/==|!=|[<>]=?/;
my $instr_re = qr/
    ^
    (?<dst_reg>$reg_re) \s+
    (?<op>$op_re) \s+
    (?<op_num>$num_re) \s+
    if \s+
    (?<cond_reg>$reg_re) \s+
    (?<cond_cmp>$cmp_re) \s+
    (?<cond_num>$num_re)
    $
/x;
sub read_input($file) {
    open my $fh, '<:encoding(UTF-8)', $file or die "$file: $!";
    my @instrs;
    while (<$fh>) {
        chomp;
        /$instr_re/ or die "failed to parse $!";
        push @instrs, {%+};
    }
    \@instrs;
}

my $op = {
    inc => '+=',
    dec => '-=',
};
sub instr_to_cmd($i) {
    sprintf(
        '$regs->{%s} %s %d if ($regs->{%s} // 0) %s %d',
        $i->{dst_reg},
        $op->{$i->{op}},
        $i->{op_num},
        $i->{cond_reg},
        $i->{cond_cmp},
        $i->{cond_num},
    );
}

sub exec_instr($regs, $i) {
    eval instr_to_cmd($i);
}


# my $instrs = read_input("test-input.txt");
my $instrs = read_input("input.txt");

say "part 1: ", solve_part_1($instrs);
say "part 2: ", solve_part_2($instrs);

sub solve_part_1($instrs) {
    my $regs = {};
    for my $i (@$instrs) {
        exec_instr($regs, $i);
    }
    scalar max values %$regs;
}

sub solve_part_2($instrs) {
    my $regs = {};
    my $max;
    for my $i (@$instrs) {
        exec_instr($regs, $i);
        my $val = $regs->{$i->{dst_reg}} // 0;
        $max = $val if !defined($max) || $val > $max;
    }
    $max;
}
