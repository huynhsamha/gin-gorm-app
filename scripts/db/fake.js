const Faker = require('fakerator');
const fs = require('fs');

const fakeUser = (fake) => {
    const first_name = fake.names.firstName();
    const last_name = fake.names.lastName();
    const username = fake.internet.userName(first_name, last_name);
    const email = fake.internet.email(first_name, last_name);
    const photoUrl = fake.internet.avatar();
    const name = first_name + ' ' + last_name;
    const password = String(username).substr(0,1) + '1';
    const github = username;
    const twitter = username;
    const website = fake.internet.url();
    const location = fake.address.streetName() + ', '  + fake.address.country()
    const title = fake.lorem.sentence()
    const aboutMe = fake.lorem.paragraph()

    return { username, email, password, name, photoUrl, github, twitter, website, location, title, aboutMe }
}

const users = [];
const locales = ['de-DE', 'en-AU', 'en-CA', 'es-ES', 'fr-FR', 'it-IT', 'sk-SK', 'sv-SE'];

const randLocale = () => {
    return Faker().random.number(locales.length-1);
}

for (let i=0;i<50;i++) {
    users.push(fakeUser(Faker(randLocale())));
}

fs.writeFileSync('./db.json', JSON.stringify(users, null, 4))
