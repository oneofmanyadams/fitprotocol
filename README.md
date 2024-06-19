# fitprotocol
A library for parsing .fit files

https://developer.garmin.com/fit/overview/

https://developer.garmin.com/fit/protocol/

ToDo

- Definition message parser
    - Read message details
        - "Record Content"
            - Read Field definitions.
            - ~Record Header (1b), Reserved(1b), Architecture(1b), etc...~
    - Find Global Fit Profile (in SDK??).
    - ~~Copy over base type definitions from SDK~~
    - ~~Header developer flag~~
    - ~~Identify Messages~~
- Start on DataRecords parsing
- Unit tests
    - CRC
    - Header
    - Data
- Parse CRC
    - Move CRC checker function out of reader and into the CRC file?
    - ~~Better standardize how CRCs are pulled out and compared?~~
        - ~~Standardize how header/data bytes are pulled out.~~
        - ~~Standardize how offsets are set.~~
    - ~~Find EOF CRC and also compare that?~~
    - ~~More structual CRC logic.~~
    - ~~Logic to easily compare Header/EOF CRC value to CRC func output?~~
    - ~~Write initial CRC func~~
- ~~Clean up how header and data reads interact.~~
    - ~~Read bytes out through separate logic.~~
    - ~~Add better erro handeling to ReadBytes method.~~
    - ~~Can the header read be defined internally~~
        - ~~Or reads just broken out to a higher level that manages the header/data/CRC split?~~
        - ~~Currently Data read only works if Header has been initiated separately.~~
- ~~Better error handling~~
    - ~~You still have println(err)'s in here dummy~~
- ~~Header parser~~
    - ~~Header auto detection? (find size function?)~~
    - ~~convert header parameters from bytes to correct data types~~
    - ~~chunk out header byte array~~
