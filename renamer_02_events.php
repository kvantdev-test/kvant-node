#!/usr/bin/php
<?php

error_reporting(0);


function getfile($d)
{
    global $l;

$ext_check[] = "go";
$ext_check[] = "sh";


$h = opendir($d);
    while($file = readdir($h))
    {
	if($file=="." || $file == "..")continue;
	if($file=="vendor")continue;
	$this_file = $d.$file;
	if(is_dir($this_file))
	getfile("$this_file/");
	else
	{
	    $t = pathinfo($this_file);
//	    if($t[extension]!="go")continue;
//print $t[extension]."\n";
	    if(!in_array($t[extension],$ext_check))continue;
	    $l[$this_file] = $this_file;
	}
    }
}


$d = __DIR__."/";
getfile($d);

$found = "MinterTeam/events-db";
$repl = basename(__DIR__)."/events-db";
//print_r($l);die;
foreach($l as $file)
{
//print $file."\n";
    $a = file_get_contents($file);
    if(strpos($a,$found)===false)continue;
    print $file."\n";
    $f = $file.".".date("renamer-Y-m-d-H-i-s");
//    file_put_contents($f,$a);
    $a = str_replace($found,$repl,$a);
    file_put_contents($file,$a);
    
}

?>
