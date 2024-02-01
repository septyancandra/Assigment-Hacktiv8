function reverseString (text) {
    const reversdString =  text.split("").reduce((acc, char) => char + acc, "");
    console.log(reversdString);
}

reverseString("Hello World and Coders");
reverseString("John Doe");
reverseString("I am a bookworm");