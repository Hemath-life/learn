# Removing System Apps on Android using ADB

Please note that removing system apps using ADB is a sensitive process and should be done with caution. Proceed at your own risk, and be aware that it may void your warranty.

## Prerequisites

- Enable "USB Debugging" on your Android device:
  - Go to "Settings."
  - Scroll down and tap on "About phone."
  - Tap "Build number" multiple times to enable "Developer Options."
  - Enable "USB debugging" in "Developer Options."

## Steps

1. **Connect Your Device to Your Computer**:

   - Use a USB cable to connect your Android device to your computer.

2. **Install ADB on Your Computer**:

   - Make sure you have ADB (Android Debug Bridge) installed on your computer. You can download the Android SDK or use a minimal ADB and Fastboot tool.
   - `sudo apt install adb`

3. **Open a Command Prompt or Terminal Window**:

   - On your computer, open a command prompt (Windows) or terminal (macOS/Linux).

4. **Check Device Connection**:

   - To ensure your Android device is recognized by ADB, run the following command:
     ```shell
     adb devices
     ```
     You should see your device listed. If not, ensure proper drivers and USB debugging.

5. **Identify the Package Name**:

   - Find the package name of the system app you want to remove using:
     ```shell
     adb shell pm list packages | grep 'keyword'
     ```
     Replace `'keyword'` with a keyword related to the app you want to remove.

6. **Uninstall the App**:

   - To uninstall the system app, use the `pm uninstall` command followed by the package name:
     ```shell
     adb shell pm list packages --user 0 // find the user 0 apps
     adb shell pm uninstall --user 0 com.android.example
     ```
     Replace `com.android.example` with the actual package name of the app.

7. **Reboot Your Device**:
   - After uninstalling the app, reboot your device to ensure proper functionality.

Remember to proceed with caution and only remove system apps that you are certain you don't need. Removing the wrong apps can affect the stability and functionality of your device.
