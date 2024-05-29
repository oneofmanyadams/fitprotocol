# fitprotocol
A library for parsing .fit files

https://developer.garmin.com/fit/overview/

https://developer.garmin.com/fit/protocol/

ToDo
- Header parser
    - Parse CRC
        - More structual CRC logic.
        - Logic to easily compare Header/EOF CRC value to CRC func output? 
        - ~~Write initial CRC func~~
    - Start on DataRecords parsing
    - ~~convert header parameters from bytes to correct data types~~
    - ~~chunk out header byte array~~
