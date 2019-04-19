$Current = [Security.Principal.WindowsIdentity]::GetCurrent()
$WindowsPrincipal = [Security.Principal.WindowsPrincipal]$Current

if (-not $WindowsPrincipal.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator))
{
    $BoundParameters = ($MyInvocation.BoundParameters.Keys | ForEach-Object { '-{0} {1}' -f  $_, $MyInvocation.BoundParameters[$_] }) -join ' '
    $Path = (Resolve-Path  $MyInvocation.InvocationName).Path
    $Parameters = $BoundParameters + ' ' + $args -join ' '
    Start-Process powershell.exe  -ArgumentList "$Path $Parameters" -verb runas -WindowStyle Hidden
    return
}

Get-Process | Where-Object { $_.ProcessName.Contains("TabTip") } | Stop-Process
