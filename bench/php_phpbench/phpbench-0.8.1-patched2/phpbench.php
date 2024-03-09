#! /usr/bin/env php
<?php
//error_reporting(0);
ignore_user_abort(TRUE);
error_reporting(E_ALL);
set_time_limit(0);
ob_implicit_flush(1);

define('PHPBENCH_VERSION', '0.8.1');
define('CSV_SEP', ',');
define('CSV_NL', "\n");
define('DEFAULT_BASE', 1000000);
define('MIN_BASE', 50);

$TESTS_DIRS = array('/usr/local/lib/phpbench/tests',
		    '/usr/local/share/phpbench/tests',
		    '/usr/lib/phpbench/tests',
		    '/usr/share/phpbench/tests',
		    '/opt/phpbench/tests',		    
		    'tests',
		    '.');

function test_start($func) {
    global $GLOBAL_TEST_FUNC;
    global $GLOBAL_TEST_START_TIME;    

    $GLOBAL_TEST_FUNC = $func;
    echo sprintf('%34s', $func) . "\t";
    flush();
    list($usec, $sec) = explode(' ', microtime());
    $GLOBAL_TEST_START_TIME = $usec + $sec;    
}

function test_end($func) {
    global $GLOBAL_TEST_FUNC;
    global $GLOBAL_TEST_START_TIME;    

    list($usec, $sec) = explode(' ', microtime());
    $now = $usec + $sec;
    if ($func !== $GLOBAL_TEST_FUNC) {
	trigger_error('Wrong func: [' . $func . '] ' .
		      'vs ' . $GLOBAL_TEST_FUNC);
	return FALSE;
    }
    if ($now < $GLOBAL_TEST_START_TIME) {
	trigger_error('Wrong func: [' . $func . '] ' .
		      'vs ' . $GLOBAL_TEST_FUNC);
	return FALSE;
    }
    $duration = $now - $GLOBAL_TEST_START_TIME;
    echo sprintf('%9.04f', $duration) . ' seconds.' . "\n";
    
    return $duration;
}

function test_regression($func) {
    trigger_error('* REGRESSION * [' . $func . ']' . "\n");
    die();
}

function do_tests($base, &$tests_list, &$results) {
    foreach ($tests_list as $test) {
	$results[$test] = call_user_func($test, $base, $results);	
    }
}

function load_test($tests_dir, &$tests_list) {
    if (($dir = @opendir($tests_dir)) === FALSE) {
	return FALSE;
    }
    $matches = array();
    while (($entry = readdir($dir)) !== FALSE) {
	if (preg_match('/^(test_.+)[.]php$/i', $entry, $matches) <= 0) {
	    continue;
	}
	$test_name = $matches[1];
	include_once($tests_dir . '/' . $entry);
	echo 'Test [' . $test_name . '] ';
	flush();
	if (!function_exists($test_name . '_enabled')) {
	    echo 'INVALID !' . "\n";	    
	    continue;
	}
	if (call_user_func($test_name . '_enabled') !== TRUE) {
	    echo 'disabled.' . "\n";
	    continue;
	}
	if (!function_exists($test_name)) {
	    echo 'BROKEN !' . "\n";
	    continue;
	}
	array_push($tests_list, $test_name);	
	echo 'enabled.' . "\n";
    }
    closedir($dir);
    
    return TRUE;	  
}

function load_tests(&$tests_dirs, &$tests_list) {
    $ret = FALSE;
    
    foreach ($tests_dirs as $tests_dir) {
	if (load_test($tests_dir, $tests_list) === TRUE) {
	    $ret = TRUE;
	}
    }
    if (count($tests_list) <= 0) {
	return FALSE;
    }
    asort($tests_list);
    
    return $ret;
}

function csv_escape($str) {
    if (strchr($str, CSV_SEP) !== FALSE) {
	return '"' . str_replace('"', '\'', $str) . '"';
    }
    return $str;
}

function export_csv($csv_file, &$results, &$percentile_times) {
    if (empty($csv_file)) {
	return TRUE;
    }
    if (($fp = fopen($csv_file, 'w')) === FALSE) {
	return FALSE;
    }
    if (fputs($fp, csv_escape('Test') . CSV_SEP . csv_escape('Time') . CSV_SEP .
	      csv_escape('Percentile') . CSV_NL)
	=== FALSE) {
	@fclose($fp);
	unlink($csv_file);
	return FALSE;
    }
    foreach ($results as $test => $time) {
	if (fputs($fp, csv_escape($test) . CSV_SEP .
		  csv_escape(sprintf('%.04f', $time)) . CSV_SEP .
		  csv_escape(sprintf('%.03f', $percentile_times[$test])) .
		  CSV_NL) === FALSE) {
	    @fclose($fp);
	    unlink($csv_file);
	    return FALSE;
	}
    }
    if (fclose($fp) === FALSE) {
	return FALSE;
    }    
    return TRUE;
}

function show_summary($base, &$results, $csv_file) {
    $total_time = 0.0;
    foreach ($results as $test => $time) {
	$total_time += $time;
    }
    if ($total_time <= 0.0) {
	die('Not enough iterations, please try with more.' . "\n");
    }
    $percentile_times = array();
    foreach ($results as $test => $time) {
	$percentile_times[$test] = $time * 100.0 / $total_time;
    }
    $score = (float) $base * 10.0 / $total_time;
    if (function_exists('php_uname')) {
	echo 'System     : ' . php_uname() . "\n";
    }
    if (function_exists('phpversion')) {
	echo 'PHP version: ' . phpversion() . "\n";
    }
    echo
      'PHPBench   : ' . PHPBENCH_VERSION . "\n" .
      'Date       : ' . date('F j, Y, g:i a') . "\n" .
      'Tests      : ' . count($results) . "\n" .
      'Iterations : ' . $base . "\n" .
      'Total time : ' . round($total_time) . ' seconds' . "\n" .
      'Score      : ' . round($score) . ' (higher is better)' . "\n";
    
    if ($csv_file !== FALSE) {
	export_csv($csv_file, $results, $percentile_times);
    }
}

function help() {
    global $TESTS_DIRS;
    
    echo
      "\n" . 'PHPBench version ' . PHPBENCH_VERSION . "\n\n" .
      '-f <file name> : Output a summary as a CSV file.' . "\n" .
      '-h             : Help.' . "\n" .
      '-i <number>    : Number of iterations (default=' . DEFAULT_BASE . ').' .
      "\n\n" .
      'Scripts are loaded from the following directories: ' . "\n";    
    foreach ($TESTS_DIRS as $tests_dir) {
	echo '  - ' . $tests_dir . "\n";
    }
    echo "\n";
}

$base = DEFAULT_BASE;
$csv_file = FALSE;

$options = getopt('f:hi:');
if (isset($options['h'])) {
    help();
    exit;
}
if (!empty($options['f'])) {
    $csv_file = $options['f'];
    if (preg_match('/[.]csv$/i', $csv_file) <= 0) {
	$csv_file .= '.csv';
    }
}
if (!empty($options['i']) && is_numeric($options['i'])) {
    $base = $options['i'];
}
if ($base < MIN_BASE) {
    die('Min iterations = ' . MIN_BASE . "\n");
}
if (empty($options)) {
    help();
}
echo 'Starting the benchmark with ' . $base . ' iterations.' . "\n\n";
$tests_list = array();
$results = array();
if (load_tests($TESTS_DIRS, $tests_list) === FALSE) {
    die('Unable to load tests');
}
echo "\n";
do_tests($base, $tests_list, $results);
echo "\n";
show_summary($base, $results, $csv_file);
echo "\n";

?>
