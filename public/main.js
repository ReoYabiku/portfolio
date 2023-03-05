window.addEventListener('load', function() {
  window.scrollBy({
    top: -40,
  });
});

window.onload = function() {
  let menu_about = document.getElementById('menu-about');
  if(menu_about != null) {
    console.log("not null");
  }
  let menu_philosophy = document.getElementById('menu-philosophy');
  let menu_biography = document.getElementById('menu-biography');
  let menu_achievement = document.getElementById('menu-achievement');
  let menu_contact = document.getElementById('menu-contact');

  let hidden = document.getElementById('hidden');

  menu_about.addEventListener('click', function() {
    hidden.clicked = false;
    console.log('clicked!');
    location.reload;
  })
  menu_philosophy.addEventListener('click', function() {
    hidden.clicked = false;
    location.reload;
  })
  menu_biography.addEventListener('click', function() {
    hidden.clicked = false;
    location.reload;
  })
  menu_achievement.addEventListener('click', function() {
    hidden.clicked = false;
    location.reload;
  })
  menu_contact.addEventListener('click', function() {
    hidden.clicked = false;
    location.reload;
  })
}
