
console.log('hello world!');


const getData = async() => {
  const res = await fetch('http://localhost:8080', {
    mode: 'cors',
  });
  const json = await res.json();
  return json.data;
}

const pTag = (value) => {
  const p = document.createElement("p"); 
  const text = document.createTextNode(value);
  p.appendChild(text);
  return p;
}

const createInfo = (e) => {
  const info = document.createElement("div");
  info.className = "info";
  const name = pTag(`Name: ${e.firstName} ${e.lastName}`);
  info.appendChild(name);

  const keys = Object.keys(e);
  for (let i = 2; i < keys.length -1; i++) {
    const key = keys[i].charAt(0).toUpperCase() + keys[i].slice(1);
    const p = pTag(`${key}: ${e[keys[i]]}`);
    info.appendChild(p);
  }
  return info;
}

const createImage = (e) => {
  const imageDiv= document.createElement("div");
  imageDiv.className = "imageDiv";
  const image = document.createElement("img");
  image.src = e.image;
  imageDiv.appendChild(image);
  return imageDiv;
}

const createPerson = (e) => {
  const div = document.createElement("div");
  div.className = "person";
  div.appendChild(createImage(e));
  div.appendChild(createInfo(e));
  return div;
}

const render = async() => {
  const data = await getData();  
  const list = document.querySelector(".list");
  data.map(e => {
    const div = createPerson(e);
    list.append(div);
  })
}



render();
