window.onload = function() {
  const menu_about = document.getElementById('menu-about');
  if(menu_about) {
    console.log("is about");
  }
  const menu_philosophy = document.getElementById('menu-philosophy');
  const menu_biography = document.getElementById('menu-biography');
  const menu_achievement = document.getElementById('menu-achievement');
  const menu_contact = document.getElementById('menu-contact');

  const menu = document.getElementById('menu');
  const shadow = document.getElementById('shadow');
  const hidden = document.getElementById('hidden');

  menu_about.addEventListener('click', function() {
    console.log("menu_about was clicked!!");
    menu.style.display = 'none';
    shadow.style.display = 'none';
    hidden.checked = false;
    location.reload();
  })
  menu_philosophy.addEventListener('click', function() {
    menu.style.display = 'none';
    shadow.style.display = 'none';
    hidden.checked = false;
    location.reload();
  })
  menu_biography.addEventListener('click', function() {
    menu.style.display = 'none';
    shadow.style.display = 'none';
    hidden.checked = false;
    location.reload();
  })
  menu_achievement.addEventListener('click', function() {
    menu.style.display = 'none';
    shadow.style.display = 'none';
    hidden.checked = false;
    location.reload();
  })
  menu_contact.addEventListener('click', function() {
    menu.style.display = 'none';
    shadow.style.display = 'none';
    hidden.checked = false;
    location.reload();
  })

  
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
