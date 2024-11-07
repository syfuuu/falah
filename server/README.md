## Features:
    * Should be able to add a mosque
        * but who can add a mosque?
            * someone should be responsible. // TODO - check with others how it should be.
            * can be more than one admin
        * how we are going to verify if the mosque is legit? with a photo for the mosque or something. // Need more ideas
        * what if responsible person is not available to update for some reason? well, we could show location based timing.
            * if location based, can also put hanafi time as well as they pray Asr bit late 
            * background jobs need to be run intervally to check if time is updated or not.
            * reminder to that person? person who are using the app could request for an update. 
            * but we have to avoid spamming.
            * push notification after some threshold.
    * Mosque details should be editable.
        * All the information
        * namaz time
        * Jummah - should be when kutbah starts
            * multiple Jumma

    * Some mosque might want to remove the mosque because they don't want to take the responsibility of updating the time or
        they don't have someone to whom they can rely on.

    * Optional login feature
        * The one who is adding the mosque must register to add
        * For users, it is optional.
            * if someone decides to register, we would need to have all his/her settings to be saved like all the mosque he has added.

    * Searching a mosque
        * should be able to search using "nearby" or actual address (maybe not full address just zipcode?)
        * link to google map if one clicks one of the mosque
        
    * Notifications to user and admin
        * we push notification to admin as a reminder everyday to update the time.
        * notifications for users when Salah time kicks in. also to admin as admin is also a user.

    * Event happening in the Masajid - optional
        * Bayans
        * Khateeb name
        * Islamic event with multiple speakers
        * any other event
        * can be set recurring
        * should be editable
        * can be deleted.


## Technical Details
    * using golang for backend. why? idk i like it and its super fast.
    * need a database. Any sql should work. Postgresql seems like good choice. 
    * for location based athan time, i couldn't find any library in golang as of now but worth to check https://github.com/ip2location/ip2location-go
        * if not, i could build a library. should be fun.
    * Database design
        * Masjid table
            * Id - string
            * Name - string
            * Details - string 
            * Address - string
            * FajrAthan - Time
            * JuhrAthan - Time
            * AsrAthan - Time
            * MagribAthan - Time
            * IshaAthan - Time
            * FajrIqamah - Time
            * JuhrIqamah - Time
            * AsrIqamah - Time
            * MagribIqamah - Time
            * IshaIqamah - Time
            * Jummah1 - Time
            * Jummah2 (Optional)
            * Jummah3 (Optional)
            * Taraweeh - Time
            * NumTaraweeh - int (Number of Taraweeh - Optional)
        * User table
            * Id - string
            * Name - string
            * Email - string
            * Password - string
            * PhoneCode - int 
            * PhoneNumber - string
            * isImam - bool
            * isAdmin - bool
        * Event table
            * Id - string
            * Type - string (jummah kutbah, bayan, etc)
            * Detail - string
            * KeyPerson - string
            * EventDate - Date
            * EventTime - Time
            * isRecurring - bool
        * Relations
            * User Id - Masjid id (ManyToMany) (affliation of the person to the masjids)
            * Event Id - Masjid id (ManyToMany) (events in the masjids)
