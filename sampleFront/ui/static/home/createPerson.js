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
