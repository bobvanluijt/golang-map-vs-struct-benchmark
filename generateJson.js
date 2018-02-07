'use strict';

// deps
var fs = require('fs');

// Counter for while loop
var n = 0;

// define objects, A = map, B = struct
var objectA = {};
var objectB = [];

// define history for objects
var history;

// minimum history, maximum history
var min = 500000;
var max = 525960;

// amount of props
var maxAmountProps = 5;

function createHistory(){

    let h = 0;
    
    let returner = {
        a: {},
        b: []
    }

    while (h < Math.floor(Math.random()*(max-min+1)+min)) {

        let string = Math.random().toString(36).substr(2, Math.floor(Math.random()*(128-1+1)+1));

        returner['a'][h] = string

        returner['b'].push({
            "timestamp": h,
            "value": string
        })
        
        h++;

    }

    return returner
}

while (n < maxAmountProps) {
   
    // get the history for this prop
    history = createHistory();

    // add the history to objectA
    objectA['prop_' + n] = history['a'];
    
    // add the history to objectB
    objectB.push({
        "property": "prop_" + n,
        "history": history['b']
    })

    history['a'] = null;
    history['b'] = null;

    n++;

    if (n == maxAmountProps){
        console.log("Id to find: ", 'prop_' + (n - 1))
    }

}

fs.writeFile("mapFile.json", JSON.stringify(objectA), function(err) {
    if (err != null){
        console.error(err);
    }
})

fs.writeFile("structFile.json", JSON.stringify(objectB), function(err) {
    if (err != null){
        console.error(err);
    }
})

console.log('done');