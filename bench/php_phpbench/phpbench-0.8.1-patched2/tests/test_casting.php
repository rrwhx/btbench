<?php
//error_reporting(0);
function test_casting($base) {
    $t = $base;
    test_start(__FUNCTION__);
    do {
	$a = 0.59;
	$b = (int) $a;
	$c = $a + $b;
	$d = 'xxxxxxxxxxxxx' . $c;
	$e = $b . $d . $a;
	$f = (float) $e;
	$g = $a . $b . $c . $d . $e . $f;
	settype($a, 'string');
    } while (--$t !== 0);

    return test_end(__FUNCTION__);
}

function test_casting_enabled() {
    return TRUE;
}

?>
