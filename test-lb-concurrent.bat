@echo off
setlocal enabledelayedexpansion

REM Test Load Balancer (Concurrent) - Swantara Gate
REM Usage: test-lb.bat [total_requests] [url]
REM Example: test-lb.bat 20 http://api.example.local:8000/

set TOTAL=%~1
set URL=%~2

if "%TOTAL%"=="" set TOTAL=20
if "%URL%"=="" set URL=http://api.example.local:8000/

echo ========================================
echo  Load Balancer Test (Concurrent) - Swantara Gate
echo ========================================
echo URL    : %URL%
echo Total  : %TOTAL% requests (concurrent)
echo ----------------------------------------

REM Create temp folder for results
set TEMP_DIR=%TEMP%\lb_test_%RANDOM%
mkdir "%TEMP_DIR%" 2>nul

REM Launch all requests concurrently using start /b
for /L %%i in (1,1,%TOTAL%) do (
    start /b "" powershell -NoProfile -Command "try { $r = Invoke-WebRequest -Uri '%URL%' -UseBasicParsing -TimeoutSec 20; [System.IO.File]::WriteAllText('%TEMP_DIR%\%%i.txt', $r.Content.Trim()) } catch { [System.IO.File]::WriteAllText('%TEMP_DIR%\%%i.txt', 'ERROR') }"
)

REM Wait for all background processes to finish
echo Waiting for all requests to complete...
:WAITLOOP
set COUNT=0
for /L %%i in (1,1,%TOTAL%) do (
    if exist "%TEMP_DIR%\%%i.txt" set /a COUNT+=1
)
if !COUNT! LSS %TOTAL% (
    timeout /t 1 /nobreak >nul
    goto WAITLOOP
)

REM Count results
set SUCCESS=0
set ERRORS=0
set SERVER_7071=0
set SERVER_7072=0

for /L %%i in (1,1,%TOTAL%) do (
    if exist "%TEMP_DIR%\%%i.txt" (
        set /p RESULT=<"%TEMP_DIR%\%%i.txt"
        if "!RESULT!"=="ERROR" (
            set /a ERRORS+=1
            echo   [%%i] ERROR
        ) else (
            set /a SUCCESS+=1
            echo   [%%i] !RESULT!
            if "!RESULT!"=="SERVER 7071" set /a SERVER_7071+=1
            if "!RESULT!"=="SERVER 7072" set /a SERVER_7072+=1
        )
    ) else (
        set /a ERRORS+=1
        echo   [%%i] TIMEOUT/NO RESPONSE
    )
)

echo.
echo ========================================
echo  HASIL
echo ========================================
echo   SERVER 7071 : %SERVER_7071%/%TOTAL%
echo   SERVER 7072 : %SERVER_7072%/%TOTAL%
echo   SUCCESS     : %SUCCESS%/%TOTAL%
echo   ERRORS      : %ERRORS%/%TOTAL%
echo ========================================

REM Cleanup
rmdir /s /q "%TEMP_DIR%" 2>nul

endlocal
