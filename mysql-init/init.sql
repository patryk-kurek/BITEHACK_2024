USE icons;

CREATE TABLE IF NOT EXISTS icons (
    id INT AUTO_INCREMENT PRIMARY KEY,
    icon_path VARCHAR(255) NOT NULL,
    description TEXT,
    tag VARCHAR(255)
);


CREATE TABLE IF NOT EXISTS tags (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS volunteers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    telephone_number VARCHAR(20) NOT NULL,
    photo VARCHAR(255),
    tags JSON NOT NULL,
    voivodeship VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS volunteer_tags (
    id INT AUTO_INCREMENT PRIMARY KEY,
    tag VARCHAR(255) NOT NULL
);

-- Insert tags
INSERT INTO tags (name) VALUES ('whatsapp');
INSERT INTO tags (name) VALUES ('ios_system');
INSERT INTO tags (name) VALUES ('android_system');
INSERT INTO tags (name) VALUES ('messenger');
INSERT INTO tags (name) VALUES ('instagram');
INSERT INTO tags (name) VALUES ('facebook');
INSERT INTO tags (name) VALUES ('google_maps');

-- Insert icons with appropriate tags
INSERT INTO icons (icon_path, description, tag)
VALUES
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-add-attachment.jpeg',
     'Use this icon to add attachments like photos or documents to your messages. Perfect for sharing files with loved ones.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-addphoto.jpeg',
     'Tap this icon to add a photo from your gallery or take a new one to send in a conversation.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-back.jpeg',
     'Use this icon to go back to the previous screen. It helps you easily navigate WhatsApp.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-calls.jpeg',
     'Tap this icon to view your recent calls or make a new one. Stay connected with friends and family.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-camera.jpeg',
     'Use this icon to open the camera. Take a picture and send it instantly to a chat.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-chat-user.jpeg',
     'This icon represents a conversation with a specific contact. Tap it to open your chat.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-chats.jpeg',
     'Tap this icon to access your list of chats. Find all your conversations here.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-community.jpeg',
     'This icon represents the WhatsApp Community section, where you can connect with groups of people.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-contact.jpeg',
     'Use this icon to view or add a contact. It makes finding and messaging people easy.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-document.jpeg',
     'Tap this icon to send a document. Ideal for sharing important files.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-examplechat.jpeg',
     'This icon shows an example of how a chat looks. Use it to learn WhatsApp messaging.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-keyboard.jpeg',
     'This icon represents the keyboard. Use it to type your messages.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-location.jpeg',
     'Tap this icon to share your location. Let your family and friends know where you are.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-phone.jpeg',
     'This icon lets you make a phone call. Stay in touch with loved ones.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-plussign.jpeg',
     'Tap this icon to access additional features like creating a new chat.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-poll.jpeg',
     'Use this icon to create and share a poll in a chat or group conversation.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-record-your-voice.jpeg',
     'This icon allows you to record a voice message. Simply speak into your phone and send.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-settings.jpeg',
     'Tap this icon to adjust your WhatsApp settings like privacy or notifications.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-sticker.jpeg',
     'This icon lets you add stickers to your messages. Make your chats fun and expressive.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-take-photo.jpeg',
     'Use this icon to take a new photo and share it directly in a conversation.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-takephoto.jpeg',
     'This icon lets you snap a quick photo and send it instantly.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-threedots.jpeg',
     'Tap this icon to view more options like settings, creating groups, or logging out.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-updates.jpeg',
     'This icon represents status updates. Check what your contacts have shared.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-video.jpeg',
     'Tap this icon to send a video or make a video call.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-icon.jpeg',
     'This icon represents the WhatsApp app. Tap it to open WhatsApp and connect with your friends and family.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/whatsapp-send.jpeg',
     'This icon lets you send your message, photo, or video instantly. Tap it when you are ready to share.',
     'whatsapp'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/googlemaps-directions.jpeg',
     'This icon represents the directions feature. Tap it to get step-by-step navigation to your destination.',
     'google_maps'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/googlemaps-keyboard.jpeg',
     'This icon shows the keyboard for entering a location or address in Google Maps.',
     'google_maps'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/googlemaps-search-button.jpeg',
     'Tap this icon to search for the location or address you have entered.',
     'google_maps'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/googlemaps-searchbar.jpeg',
     'This icon represents the search bar. Use it to type the name of a place, address, or business.',
     'google_maps'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/googlemaps-start.jpeg',
     'Tap this icon to start navigation. Follow the instructions to reach your destination safely.',
     'google_maps'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/googlemaps-icon.jpeg',
     'This icon represents the Google Maps app. Tap it to open Google Maps and begin exploring.',
     'google_maps'),
     ('https://bitehack.s3.eu-north-1.amazonaws.com/ios_system-display-and-brightness.jpeg',
     'This icon lets you adjust the display brightness on your iPhone. Use it to make the screen easier to see, especially in low light.',
     'ios_system'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/ios_system-font_size_slider.jpeg',
     'This icon allows you to change the font size on your iPhone. Use it to make text bigger and easier to read.',
     'ios_system'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/ios_system-settings_icon.jpeg',
     'This is the Settings icon. Tap it to access all your iPhone settings, including Wi-Fi, display, and privacy options.',
     'ios_system'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/ios_system-text_size.jpeg',
     'This icon allows you to adjust the text size. Make text larger or smaller depending on what feels comfortable to read.',
     'ios_system'),
      ('https://bitehack.s3.eu-north-1.amazonaws.com/telephone_recents_system_ios.jpeg',
     'This icon shows your recent calls. Tap it to view who you called or missed.',
     'ios_system'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/telephone_new_contact_done_system_ios.jpeg',
     'This icon confirms that you have successfully saved a new contact.',
     'ios_system'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/telephone_new_contact_system_ios.jpeg',
     'This icon allows you to add a new contact to your phone. Fill in the details and tap save.',
     'ios_system'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/telephone_icon_system_ios.jpeg',
     'This is the main Telephone icon. Tap it to make calls, view contacts, or access recent calls.',
     'ios_system'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/telephone_keyboard_system_ios.jpeg',
     'This icon opens the keypad to manually dial a phone number.',
     'ios_system'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/telephone_keypad_system_ios.jpeg',
     'This icon represents the number keypad. Use it to type in a phone number to call.',
     'ios_system'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/telephone_favourites_system_ios.jpeg',
     'This icon takes you to your favorite contacts. Save frequently called people here for quick access.',
     'ios_system'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/new_contact_telephone_new_contact_typing_telephone_system_ios.jpeg',
     'This icon shows the process of typing a new contact. Enter details here and save them for future calls.',
     'ios_system'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/new_contact_telephone_phone_number_system_ios.jpeg',
     'This icon represents adding a phone number to a new contact. Tap to input the number.',
     'ios_system'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/contacts_icon_system_ios.jpeg',
     'This is the Contacts icon. Tap it to access or search for saved phone numbers.',
     'ios_system'),
    ('https://bitehack.s3.eu-north-1.amazonaws.com/contacts_system_ios.jpeg',
     'This icon represents your phonebook. Use it to find, add, or edit your contacts.',
     'ios_system');

INSERT INTO volunteers (name, surname, email, telephone_number, photo, tags, voivodeship) VALUES
('Karolina', 'Sadowska', 'karolina.sadowska@example.com', '500123456', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["SPRZATANIE", "OGROD"]', 'LUBUSKIE'),
('Dominik', 'Kaczmarek', 'dominik.kaczmarek@example.com', '501234567', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["POSILKI", "ZAKUPY"]', 'MAZOWIECKIE'),
('Natalia', 'Witkowska', 'natalia.witkowska@example.com', '502345678', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["LEKI", "OPIEKA", "ROZMOWA"]', 'SWIETOKRZYSKIE'),
('Bartosz', 'Lis', 'bartosz.lis@example.com', '503456789', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["TRANSPORT", "NAPRAWA"]', 'WARMIANSKO_MAZURSKIE'),
('Monika', 'Król', 'monika.krol@example.com', '504567890', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["TECHNOLOGIA", "PISANIE"]', 'PODLASKIE'),
('Adam', 'Jaworski', 'adam.jaworski@example.com', '505678901', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["SPACER", "OGROD", "NAPRAWA"]', 'OPOLSKIE'),
('Joanna', 'Gorska', 'joanna.gorska@example.com', '506789012', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["ROZMOWA", "LEKI"]', 'ZACHODNIOPOMORSKIE'),
('Pawel', 'Czarnecki', 'pawel.czarnecki@example.com', '507890123', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["ZAKUPY", "SPRZATANIE"]', 'SLASKIE'),
('Izabela', 'Wojciechowska', 'izabela.wojciechowska@example.com', '508901234', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["OPIEKA", "SPACER"]', 'WIELKOPOLSKIE'),
('Krzysztof', 'Kowal', 'krzysztof.kowal@example.com', '509012345', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["NAPRAWA", "TECHNOLOGIA", "TRANSPORT"]', 'DOLNOSLASKIE'),
('Sylwia', 'Michalska', 'sylwia.michalska@example.com', '510123456', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["POSILKI", "PISANIE"]', 'KUJAWSKO_POMORSKIE'),
('Tadeusz', 'Olszewski', 'tadeusz.olszewski@example.com', '511234567', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["OGROD", "NAPRAWA"]', 'MALOPOLSKIE'),
('Ewa', 'Piotrowska', 'ewa.piotrowska@example.com', '512345678', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["SPRZATANIE", "ROZMOWA", "OPIEKA"]', 'LUBELSKIE'),
('Grzegorz', 'Zajac', 'grzegorz.zajac@example.com', '513456789', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["TECHNOLOGIA", "LEKI"]', 'POMORSKIE'),
('Alicja', 'Stepien', 'alicja.stepien@example.com', '514567890', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["TRANSPORT", "ZAKUPY", "POSILKI"]', 'PODKARPACIE'),
('Mateusz', 'Kurek', 'mateusz.kurek@example.com', '515678901', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["ROZMOWA", "SPACER"]', 'SWIETOKRZYSKIE'),
('Zofia', 'Majewska', 'zofia.majewska@example.com', '516789012', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["SPRZATANIE", "OGROD", "POSILKI"]', 'WIELKOPOLSKIE'),
('Mariusz', 'Jankowski', 'mariusz.jankowski@example.com', '517890123', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["NAPRAWA", "TRANSPORT"]', 'MAZOWIECKIE'),
('Patrycja', 'Nowinska', 'patrycja.nowinska@example.com', '518901234', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["PISANIE", "TECHNOLOGIA"]', 'SLASKIE'),
('Jakub', 'Kowalczyk', 'jakub.kowalczyk@example.com', '519012345', 'https://bitehack.s3.eu-north-1.amazonaws.com/image.jpeg', '["LEKI", "ZAKUPY"]', 'DOLNOSLASKIE');


INSERT INTO volunteer_tags (tag) VALUES ('ZAKUPY');
INSERT INTO volunteer_tags (tag) VALUES ('LEKI');
INSERT INTO volunteer_tags (tag) VALUES ('SPRZATANIE');
INSERT INTO volunteer_tags (tag) VALUES ('SPACER');
INSERT INTO volunteer_tags (tag) VALUES ('TRANSPORT');
INSERT INTO volunteer_tags (tag) VALUES ('POSILKI');
INSERT INTO volunteer_tags (tag) VALUES ('OPIEKA');
INSERT INTO volunteer_tags (tag) VALUES ('ROZMOWA');
INSERT INTO volunteer_tags (tag) VALUES ('NAPRAWA');
INSERT INTO volunteer_tags (tag) VALUES ('TECHNOLOGIA');
INSERT INTO volunteer_tags (tag) VALUES ('OGROD');
INSERT INTO volunteer_tags (tag) VALUES ('PISANIE');