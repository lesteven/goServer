
console.log('hello world!');


const getData = async() => {
  const res = await fetch('http://localhost:8080', {
    mode: 'cors',
  });
  const json = await res.json();
  console.log(json);
}

getData();
