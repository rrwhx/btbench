<?php

function tofr_5(&$a, &$b) {
    return $a . $b;
}

function tofr_4(&$a, &$b) {
    return tofr_5($a, $b);
}

function tofr_3(&$a, &$b) {
    return tofr_4($a, $b);
}

function tofr_2(&$a, &$b) {
    return tofr_3($a, $b);
}

function tofr_1(&$a, $b) {
    return tofr_2($a, $b);
}

function test_ordered_functions_references($base) {
    $t = $base;
    test_start(__FUNCTION__);
    $z = 42;
    $y = -1;
    $x = 42.67;
    $w = 'xxx';
    $v = NULL;
    do {
	$a = tofr_1($z, $y);
	$b = tofr_1($a, $x);
	$c = tofr_1($b, $w);
	$d = tofr_1($c, $v);
	$e = tofr_1($a, tofr_1($b, tofr_1($c, tofr_1($d, $a))));
    } while (--$t !== 0);

    return test_end(__FUNCTION__);
}

function test_ordered_functions_references_enabled() {
    return TRUE;
}

?>
