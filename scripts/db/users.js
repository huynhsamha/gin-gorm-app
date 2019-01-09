const users = require('./users.json');

const async = require('async');
const axios = require('axios');
const qs = require('querystring');

const api = 'http://localhost:8001/api/fakeDB/users'
const Authorization = 'uher0823y023s'

function genOne(user) {
  axios.post(api, qs.stringify(user), { headers: { Authorization } })
    .then(res => res.data).then(data => console.log(data)).catch(err => console.log(err))
}

function genAll() {
  async.eachSeries(users, (user, cb) => {
    console.log(user);

    axios.post(api, qs.stringify(user), { headers: { Authorization } })
      .then(() => cb()).catch(err => cb(err))

  }, (err) => {
    if (err) console.log(err);
    else console.log('OK');

    process.exit(0);
  })
}

// genOne(users[2]);
genAll();
