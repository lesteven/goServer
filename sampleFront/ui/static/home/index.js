const getData = async() => {
  const res = await fetch('http://localhost:8080', {
    mode: 'cors',
  });
  const json = await res.json();
  return json.data;
}

const render = (data) => {
  const list = document.querySelector(".list");
  data.map(e => {
    const div = createPerson(e);
    list.append(div);
  })
}

const clear = () => {
  const list = document.querySelector(".list");
  while (list.firstChild) {
    list.removeChild(list.firstChild);
  }
}

const inSearch = (data, searchVal) => {
  const filtered = data.filter(e => {
    const name = `${e.firstName} ${e.lastName}`.toLowerCase();
    return name.includes(searchVal)
  })
  return filtered;
}

const filter = (data) => (e) => {
  const searchVal = e.target.value.trim();
  const list = document.querySelector(".list");
  // prevent filtering when searchVal is empty
  if (searchVal == "" && data.length == list.children.length) {
    return;
  }
  clear();
  const filteredData = inSearch(data, searchVal);
  render(filteredData);
};

const init  = async() => {
  const data = await getData();  
  render(data);

  const input = document.querySelector(".search");
  input.addEventListener("keyup", filter(data));
}

init();
