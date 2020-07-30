# pfsense Config Tool

This simple tool helps you quickly and efficiently change the inet devices in your config file in order to prep it for use on new hardware.

## Usage

The binary executables for Linux, macOS and Windows are already included in the `bin/` directory. Simply run the tool with the config you want to use and where you want the altered config saved. The original config will not be altered. 

Eg;

```
./bin/pfsenseconftool_linux -in config.xml -out other.xml 
```

```
pfsenseconftool_win.exe -in config.xml -out other.xml 
```

Follow the prompts to confirm your existing interfaces and enter in the new interfaces. It's really that simple.

## Building

If you would like to build it yourself, builds have been tested with `go1.14.1` but should would on nearly any modern version of Go.

Use `make` to build Linux, macOS, and Windows versions of the tool in the `bin/` directory or use `build-linux`, `build-macos`, or `build-windows` to build the individual versions.

### Note

This program and related code is provided "as is" and any express or implied warranties, including the implied warranties of merchantability and fitness for a particular purpose are disclaimed. In no event shall Computer Assistance or contributors be liable for any direct, indirect, incidental, special, exemplary, or consequential damages (including, but not limited to, procurement of substitute goods or services; loss of use, data, or profits; or business interruption) sustained by you or a third party, however caused and on any theory of liability, whether in contract, strict liability, or tort arising in any way out of the use of this sample code, even if advised of the possibility of such damage.