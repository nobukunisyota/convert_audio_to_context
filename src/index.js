import { username } from "./setenv.js";

async function getData() {
    const response = await fetch(`https://zenn.dev/api/articles?username=${username}`);
    const data = await response.json();
    return data;
}
try {
    getData()
    .then(function(data) {
        for (let i = 0; i < data.length; i++) {
            console.log(data[i].name);
        }
    })
}
catch (error) {
    console.log(error);
}
finally {
    console.log('This is always executed');
}
