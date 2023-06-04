const menu = document.getElementById('menu');
const shadow = document.getElementById('shadow');
const hidden = document.getElementById('hidden');

function main() {
  const domain = "https://portfolio-go.yabi-works.com";

  optimizeHamburgerMenu();
  fetchTemplateCardContents(domain+"/biographies", '#biography-template-wrapper');
  fetchTemplateCardContents(domain+"/achievements", '#achievement-template-wrapper');
  openAccordion();
}

function optimizeHamburgerMenu() {
  const menu_ids = [
    '#menu-about',
    '#menu-philosophy',
    '#menu-biography',
    '#menu-achievement',
    '#menu-contact'
  ];
  menu_ids.forEach(menu_id => {
    document.querySelector(menu_id).addEventListener('click', closeHamburgerMenu);
  });
}

function closeHamburgerMenu() {
  menu.style.display = 'none';
  shadow.style.display = 'none';
  hidden.checked = false;
  location.reload();
}

function fetchTemplateCardContents(url = "", wrapper_id = "") {
  var request = new XMLHttpRequest();
  request.open('GET', url, false);
  request.send(null);
  var biography_contents = JSON.parse(request.responseText);
  
  if('content' in document.createElement('template')) {
    var biography_wrapper = document.querySelector(wrapper_id);
    var template_card = document.querySelector('#template-card');
    
    biography_contents.forEach(biography_content => {
      var clone = template_card.content.cloneNode(true);
      var image = clone.querySelector('img');
      var summary = clone.querySelector('summary');
      var p = clone.querySelectorAll('p');
      
      image.src         = biography_content.src;
      image.alt         = biography_content.alt;
      summary.innerText = biography_content.summary;
      p[0].textContent  = biography_content.title;
      p[1].innerText    = biography_content.description;
      
      biography_wrapper.appendChild(clone);
    })
  }
}

function openAccordion() {
  const details_list = document.getElementsByTagName('details');
  if(window.outerWidth > 900) {
    for(const details of details_list) {
      details.open = true;
      details.firstElementChild.style.listStyle = 'none';
      details.addEventListener('toggle', function() {
        details.open = true;
      });
    }
  }
}

window.onload = main();