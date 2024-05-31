# fitprotocol
A library for parsing .fit files

https://developer.garmin.com/fit/overview/

https://developer.garmin.com/fit/protocol/

ToDo
- Header parser
    - Header auto detection? (find size function?)
    - ~~convert header parameters from bytes to correct data types~~
    - ~~chunk out header byte array~~
- Start on DataRecords parsing
- Parse CRC
    - Better standardize how CRCs are pulled out and compared?
    - ~~Find EOF CRC and also compare that?~~
    - ~~More structual CRC logic.~~
    - ~~Logic to easily compare Header/EOF CRC value to CRC func output?~~
    - ~~Write initial CRC func~~

