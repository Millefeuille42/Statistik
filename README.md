# Statistik
A stat scanner for the Synthetik game. (Beta V1.0)

## How to use
### Configuration
You will find a file named `statistik_conf.json` this file must exist and must be properly filled.<br>
The different options are:

-   `tempFilePath`: (`./.temp.json`) The temp file path where the temporary data is stored in order to leave your ram free during your run. 
    You can leave this setting as is. 
    You just have to make sure the software has Read-Write access to this path.
    
-   `statsPath`: (`Statistics.sav`) The path for the `Statistics.sav` file,
    usually located at `C:\Users\YOU\.AppData\Local\Synthetik\Statistics.sav` under Windows.
    
-   `statsOutputDirectory`: (`./`) The generated stats output directory,
    make the sure the targeted folder exists and has Read-Write access. It works with ou without a trailing slash (/)

### Starting the app
Start the app by starting it like any other app, it will then open a terminal for you to enter commands.

#### The Commands
-   `start`: Start collecting data, use this command before starting your run.
    
-   `stop`: Stop collecting data and generate the run data, 
    use this command after your run is over (once you're back to the lobby).
    
-   `exit`: Exit the program

##Known Bugs
- Backslashes (&#92;) Are not currently supported however,
even in Windows styled path you can use forward slashes (/). If you absolutely want backslashes,
you'll have to double them. (&#92;&#92;)
  
- The scanner might crash because of your temp file for unknown reasons,
  sometimes because of bad JSON formatting. I'm currently working on this.
  In the meantime you can test jour temp JSON file on an online tester and fix the formatting yourself (sorry T_T).
  In this case, you can then launch the command stop, again, to create your stat file.
## Roadmap
- Proper Graphical Interface
- Plots and Graphics
- Overall Stats
