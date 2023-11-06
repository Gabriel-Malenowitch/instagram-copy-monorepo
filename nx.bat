set path_to_root=%~dp0
@REM WHERE node >nul 2>nul
@REM IF %ERRORLEVEL% NEQ 0 (ECHO Nx requires NodeJS to be available. To install NodeJS and NPM, see: https://nodejs.org/en/download/ .; EXIT 1)
@REM WHERE npm >nul 2>nul
@REM IF %ERRORLEVEL% NEQ 0 (ECHO Nx requires npm to be available. To install NodeJS and NPM, see: https://nodejs.org/en/download/ .; EXIT 1)
node %path_to_root%\.nx\nxw.js %*