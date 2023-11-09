$runtime = $MyInvocation.MyCommand.Path
$bin = $runtime + "/bin"

$env:PATH = $env:PATH + ":" + $bin
