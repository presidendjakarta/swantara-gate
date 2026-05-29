# Test Load Balancer - Swantara Gate
# Usage: .\test-lb.ps1 [-Total 20] [-Url "http://api.example.local:8000/"]
# cd x:\laragon\go-apps\swantara-gate ; powershell -ExecutionPolicy Bypass -File .\test-lb.ps1 -Total 10


param(
    [int]$Total = 20,
    [string]$Url = "http://api.example.local:8000/"
)

Write-Host "========================================" -ForegroundColor Cyan
Write-Host " Load Balancer Test - Swantara Gate" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "URL    : $Url"
Write-Host "Total  : $Total requests"
Write-Host "----------------------------------------"

$results = @{}
$errors = 0

for ($i = 1; $i -le $Total; $i++) {
    try {
        $response = Invoke-WebRequest -Uri $Url -UseBasicParsing -TimeoutSec 5
        $body = $response.Content.Trim()
        
        if ($results.ContainsKey($body)) {
            $results[$body]++
        } else {
            $results[$body] = 1
        }
        
        Write-Host "  [$i] $body" -ForegroundColor Green
    }
    catch {
        $errors++
        Write-Host "  [$i] ERROR: $($_.Exception.Message)" -ForegroundColor Red
    }
}

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host " HASIL" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan

foreach ($key in $results.Keys | Sort-Object) {
    $count = $results[$key]
    $pct = [math]::Round(($count / $Total) * 100, 1)
    $bar = "#" * [math]::Floor($pct / 2)
    Write-Host "  $key : $count/$Total ($pct%) $bar" -ForegroundColor Yellow
}

if ($errors -gt 0) {
    Write-Host "  ERRORS : $errors/$Total" -ForegroundColor Red
}

Write-Host "========================================" -ForegroundColor Cyan
