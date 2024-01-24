function buttonClick() {
    fetch('https://zipcloud.ibsnet.co.jp/api/search?zipcode=7830060')
    .then(response => response.json())
    .then((json) => {
        let hoge = json;
        console.log(hoge.results[0]);
        let value = Object.values(hoge.results[0]);
        console.log(value);
    });
}
