# fitprotocol
A library for parsing .fit files

https://developer.garmin.com/fit/overview/

https://developer.garmin.com/fit/protocol/

ToDo
- Clean up how header and data reads interact.
    - Can the header read be defined internally
        - Or reads just broken out to a higher level that manages the header/data/CRC split?
        - Currently Data read only works if Header has been initiated separately.
- Start on DataRecords parsing
- Unit tests
    - CRC
    - Header
    - Data
- Header parser
    - Header auto detection? (find size function?)
    - ~~convert header parameters from bytes to correct data types~~
    - ~~chunk out header byte array~~
- Parse CRC
    - Better standardize how CRCs are pulled out and compared?
    - ~~Find EOF CRC and also compare that?~~
    - ~~More structual CRC logic.~~
    - ~~Logic to easily compare Header/EOF CRC value to CRC func output?~~
    - ~~Write initial CRC func~~
- ~~Better error handling~~
    - ~~You still have println(err)'s in here dummy~~
