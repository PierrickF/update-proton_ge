This program compares the latest release of Proton-GE to the most recent
version installed on your system, and will update it if a newer version exists.

It is a reworked and improved version of the script given on the [Proton-GE repository](https://github.com/GloriousEggroll/proton-ge-custom).

Make sure to read the documentation given for Proton-GE.

---

# Installation

## 1. Download with Git (preferred)

1. Set your current working directory to a place you deem approriate for downloads:

`cd ~/downloads/`
2. Clone this repository:

`git clone https://github.com/PierrickF/update-proton_ge.git`
3. Set your current working directory to the newly created directory:

`cd update-proton_ge`

## 1. Download the .zip file (optionnal)

1. Click the green `Code` button.
2. Click `Download ZIP`.
3. Set your current working directory to your default downloads directory:

`cd ~/downloads/`
4. Unzip the file:

`unzip update-proton_ge-main.zip`
5. Set your current working directory to the newly created directory:

`cd update-proton_ge-main`

## 2. Copy the program to the right directory

1. List the directories found in your `PATH` environment variable:

`echo $PATH`
2. Choose one of those directories and copy the program file to it:

`cp update-proton_ge <directory_in_your_path>`
3. Give the `update_proton-ge` file appropriate permissions:

`chmod +x <directory_in_your_path>/update-proton_ge`

---

# Usage

You can now run the program from anywhere in the terminal:

`update-proton_ge`


You can also refer it by its name `update-proton_ge` in other scripts, or
in scheduled timers or events, for instance.
