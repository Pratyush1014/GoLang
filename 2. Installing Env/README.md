# Environment Setup

## Steps
To get going with the development in golang follow the below steps:

1. Install windows subsystem for linux.
    fire up your powershell terminal and type in the following commands
    ```
    wsl --install
    ```
2. Reboot the PC.
3. Create credentials for UNIX subsystem.
4. Check for potential updates
    ```
    sudo apt update
    ```
5. Install make
    ```
    sudo apt install make
    ```
    To verify check make's version
    ```
    make --version
    ```
    To open wsl terminal look for ubuntu in finder or startmenu.
6. Install go lang
    ```
    sudo snap install go --classic
    ```
    Check version
    ```
    go version
    ```
7. Install sqlc, usefull in generting golang codes
    ```
    sudo snap install sqlc
    ```
8. Configure vscode environment
    Install wsl and go lang plugins, aids in go lang developemnt and interacting with the linux subsystem

    Note :
        APT (Advanced Package Tool) and Snap are both package management systems for Linux-based operating systems. APT is the traditional package manager for Debian and Ubuntu-based distributions, while Snap is another, newer package management system developed by Canonical that works across a range of Linux distributions.
9. Install Docker
    Download Docker but don't just install it, might result in mounting issues which stops go from working
    Navigate to the directory where docker installer recides
    Run the below comman in elevated powershell 
    ```
    Start-Process "Docker Desktop Installer.exe" -Verb RunAs -Wait -ArgumentList "install --installation-dir=C:\Docker\"
    ```

    Restart your PC, open up docker head to settings and check the following
    - Use the WSL2 based engine and check the add *docker names to host machine, allows communication between container and host machine.
10. Final check 
    ```
    docker ps
    go version
    ```
    If you get the expected outputs voila your environments setup.
