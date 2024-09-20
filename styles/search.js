 
  let data = []; // Define data globally

  function processForm() {
      const forms = document.querySelectorAll('.box-form');
      
      data = Array.from(forms).map(form => {
          const inputs = form.querySelectorAll('input');
          const button = form.querySelector('.box');
          
          const inputValues = {};
          inputs.forEach(input => {
              inputValues[input.name] = input.value;
          });
          
          return {
              name: inputValues['Name'],
              members: inputValues['Members'].split(',').filter(member => member),
              year: inputValues['Year'],
              album: inputValues['Album'],
              locations: inputValues['Locations'].split('/').filter(location => location),
              button: button
          };
      });
  }
  
  
  function filterSuggestions() {
      const input = document.getElementById('search-bar').value.toLowerCase();
  
      // If the input is empty, show all buttons
      if (!input) {
          data.forEach(item => {
              item.button.style.display = 'flex'; // Show all buttons
              resetMembersDisplay(item);
          });
          return;
      }
  
      // Filter data based on input
      data.forEach(item => {
          const nameMatches = item.name.toLowerCase().includes(input);
          const yearMatches = item.year.toLowerCase().includes(input);
          const albumMatches = item.album.toLowerCase().includes(input);
          const matchingMembers = item.members.filter(member => member.toLowerCase().includes(input));
          const memberMatches = matchingMembers.length > 0;
          const matchingLocations = item.locations.filter(location => location.toLowerCase().includes(input));
          const locationMatches = matchingLocations.length > 0;
  
          // Show or hide buttons based on match
          if (nameMatches || yearMatches || albumMatches || memberMatches || locationMatches) {
              item.button.style.display = 'flex';
              if (yearMatches) {
                  updateYearDisplay(item.button, item.year); // Update the Members div with matching Year
              } else if (albumMatches) {
                  updateAlbumDisplay(item.button, item.album); // Update the Members div with matching Album
              } else if (locationMatches) {
                  updateLocationsDisplay(item.button, matchingLocations); // Update the Members div with matching Locations
              } else if (memberMatches) {
                  updateMembersDisplay(item.button, matchingMembers); // Update the Members div with matching Members
              } else {
                  resetMembersDisplay(item); // Reset to show all members
              }
          } else {
              item.button.style.display = 'none';
          }
      });
  }
  
  
  
  function updateMembersDisplay(button, matchingMembers) {
      const membersDiv = button.querySelector('.Members');
      membersDiv.innerHTML = ''; // Clear current members
  
      // Create new divs for matching members and apply color change
      matchingMembers.forEach(member => {
          const memberDiv = document.createElement('div');
          memberDiv.textContent = `[${member}]`;
          memberDiv.style.color = '#A9D8D9'; // Change color to desired color
          membersDiv.appendChild(memberDiv);
      });
  }
  
  function updateLocationsDisplay(button, matchingLocations) {
      const membersDiv = button.querySelector('.Members');
      membersDiv.innerHTML = ''; // Clear current members
  
      // Create new divs for matching locations and apply color change
      matchingLocations.forEach(location => {
          const locationDiv = document.createElement('div');
          locationDiv.textContent = `[${location}]`;
          locationDiv.style.color = '#A9D8D9'; // Change color to desired color
          membersDiv.appendChild(locationDiv);
      });
  }
  
  function resetMembersDisplay(item) {
      const membersDiv = item.button.querySelector('.Members');
      membersDiv.innerHTML = ''; // Clear current members
      item.members.forEach(member => {
          const memberDiv = document.createElement('div');
          memberDiv.textContent = `[${member}]`;
          memberDiv.style.color = ''; // Reset color
          membersDiv.appendChild(memberDiv);
      });
  }
  function updateYearDisplay(button, year) {
      const membersDiv = button.querySelector('.Members');
  
      // Create a new div for the matching year and apply color change
      const yearDiv = document.createElement('div');
      membersDiv.appendChild(yearDiv);
  }
  
  function updateAlbumDisplay(button, album) {
      const membersDiv = button.querySelector('.Members');
      membersDiv.innerHTML = ''; // Clear current members
  
      // Create a new div for the matching album and apply color change
      const albumDiv = document.createElement('div');
      albumDiv.textContent = `First Album: [${album}]`;
      albumDiv.style.color = '#A9D8D9'; // Change color to desired color
      membersDiv.appendChild(albumDiv);
  }
  
  function updateMembersDisplay(button, matchingMembers) {
      const membersDiv = button.querySelector('.Members');
      membersDiv.innerHTML = ''; // Clear current members
  
      // Create new divs for matching members and apply color change
      matchingMembers.forEach(member => {
          const memberDiv = document.createElement('div');
          memberDiv.textContent = `[${member}]`;
          memberDiv.style.color = '#A9D8D9'; // Change color to desired color
          membersDiv.appendChild(memberDiv);
      });
  }
  
  function updateLocationsDisplay(button, matchingLocations) {
      const membersDiv = button.querySelector('.Members');
      membersDiv.innerHTML = ''; // Clear current members
  
      // Create new divs for matching locations and apply color change
      matchingLocations.forEach(location => {
          const locationDiv = document.createElement('div');
          locationDiv.textContent = `[${location}]`;
          locationDiv.style.color = '#A9D8D9'; // Change color to desired color
          membersDiv.appendChild(locationDiv);
      });
  }
  
  function resetMembersDisplay(item) {
      const membersDiv = item.button.querySelector('.Members');
      membersDiv.innerHTML = ''; // Clear current members
      item.members.forEach(member => {
          const memberDiv = document.createElement('div');
          memberDiv.textContent = `[${member}]`;
          memberDiv.style.color = ''; // Reset color
          membersDiv.appendChild(memberDiv);
      });
  }
  
  // Initialize the data array on page load
  window.onload = processForm;