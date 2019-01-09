const users = require('./db.json');

const async = require('async');
const axios = require('axios');
const qs = require('querystring');

axios.post('http://localhost:8001/api/fakeDB/users', qs.stringify(users[0]))
    .then(res => console.log(res)).catch(err => console.log(err))

// async.eachSeries(users, (user, cb) => {
//     console.log(user);

//     axios.post('http://localhost:8001/api/fakeDB/users', qs.stringify(user))
//     .then(() => cb()).catch(err => cb(err))

// }, (err) => {
//     if (err) console.log(err);
//     else console.log('OK');

//     process.exit(0);
// })
