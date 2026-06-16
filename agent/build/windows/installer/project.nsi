Unicode true

# --- 1. Header & Modern UI ---
!include "MUI2.nsh"
!include "wails_tools.nsh"

# --- 2. Configuration & Branding ---
# Prevents redefinition errors by checking if Wails already set them
!ifndef INFO_COMPANYNAME
    !define INFO_COMPANYNAME "Fehmi"
!endif
!ifndef INFO_PRODUCTNAME
    !define INFO_PRODUCTNAME "Endpoint agent"
!endif
!ifndef INFO_PRODUCTVERSION
    !define INFO_PRODUCTVERSION "1.0"
!endif


!define MUI_WELCOMEPAGE_TITLE "Welcome to the Fehmi Endpoint Agent Setup Wizard"
#!define MUI_WELCOMEPAGE_TEXT "Setup will guide you through the installation of ${INFO_PRODUCTNAME}.\r\n\r\nClick Next to continue."

# Meta-Data
VIProductVersion "${INFO_PRODUCTVERSION}.0"
VIFileVersion    "${INFO_PRODUCTVERSION}.0"
VIAddVersionKey "CompanyName"     "${INFO_COMPANYNAME}"
VIAddVersionKey "FileDescription" "${INFO_PRODUCTNAME} Installer"
VIAddVersionKey "ProductVersion"  "${INFO_PRODUCTVERSION}"
VIAddVersionKey "FileVersion"     "${INFO_PRODUCTVERSION}"
VIAddVersionKey "LegalCopyright"  "© 2026 Fehmi Corporation"
VIAddVersionKey "ProductName"     "${INFO_PRODUCTNAME}"

ManifestDPIAware true

# --- 3. Interface Settings ---
!define MUI_ICON "..\icon.ico"
!define MUI_UNICON "..\icon.ico"
!define MUI_FINISHPAGE_NOAUTOCLOSE
!define MUI_ABORTWARNING

# License Page Settings
!define MUI_LICENSEPAGE_CHECKBOX
!define MUI_LICENSEDATA "${__FILEDIR__}\license.txt"

# --- 4. Page Order ---
!insertmacro MUI_PAGE_WELCOME
!insertmacro MUI_PAGE_LICENSE "${MUI_LICENSEDATA}"
!insertmacro MUI_PAGE_DIRECTORY
!insertmacro MUI_PAGE_INSTFILES
!insertmacro MUI_PAGE_FINISH

!insertmacro MUI_UNPAGE_INSTFILES
!insertmacro MUI_LANGUAGE "English"

# --- 5. Installer Settings ---
Name "Fehmi Endpoint Agent"
OutFile "..\..\bin\${INFO_PROJECTNAME}-${ARCH}-installer.exe"
InstallDir "$PROGRAMFILES64\${INFO_COMPANYNAME}\agent"
ShowInstDetails show

Function .onInit
   !insertmacro wails.checkArchitecture
FunctionEnd

# --- 6. Installation Sections ---
Section
    !insertmacro wails.setShellContext
    !insertmacro wails.webview2runtime

    SetOutPath $INSTDIR
    !insertmacro wails.files

    # Prepare ProgramData for the Go app to populate later
    CreateDirectory "$PROGRAMDATA\Fehmi\endpoint agent"
    
    CreateShortcut "$SMPROGRAMS\${INFO_PRODUCTNAME}.lnk" "$INSTDIR\${PRODUCT_EXECUTABLE}"
    CreateShortCut "$DESKTOP\${INFO_PRODUCTNAME}.lnk" "$INSTDIR\${PRODUCT_EXECUTABLE}"

    !insertmacro wails.associateFiles
    !insertmacro wails.associateCustomProtocols
    !insertmacro wails.writeUninstaller
SectionEnd

Section "uninstall"
    !insertmacro wails.setShellContext
    RMDir /r "$AppData\${PRODUCT_EXECUTABLE}" 
    RMDir /r $INSTDIR
    # Cleanup the config folder
    RMDir /r "$PROGRAMDATA\Fehmi"
    
    Delete "$SMPROGRAMS\${INFO_PRODUCTNAME}.lnk"
    Delete "$DESKTOP\${INFO_PRODUCTNAME}.lnk"
    
    !insertmacro wails.unassociateFiles
    !insertmacro wails.unassociateCustomProtocols
    !insertmacro wails.deleteUninstaller
SectionEnd