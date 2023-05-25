const menu = document.getElementById('menu');
const shadow = document.getElementById('shadow');
const hidden = document.getElementById('hidden');

function main() {
  const menu_ids = [
    '#menu-about',
    '#menu-philosophy',
    '#menu-biography',
    '#menu-achievement',
    '#menu-contact'
  ];
  menu_ids.forEach(menu_id => {
    document.querySelector(menu_id).addEventListener('click', closeHumburgerMenu);
  });

  const details_list = document.getElementsByTagName('details');
  if(window.outerWidth > 1300) {
    for(const details of details_list) {
      details.open = true;
      details.firstElementChild.style.listStyle = 'none';
      details.addEventListener('toggle', function() {
        details.open = true;
      });
    }
  }
}

function closeHumburgerMenu() {
  menu.style.display = 'none';
  shadow.style.display = 'none';
  hidden.checked = false;
  location.reload();
}

window.onload = main();