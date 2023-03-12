# chatgptdash

## What is it
This is a command line tool to access Chat GPT from command line.  This is build using Go.

## How to Install

Run the following commands

```
brew tap dash2701/dash2701
```

```
brew install dash2701/dash2701/chatgptdash
```

## How to use it 

To use this from command line, we would need to generate an token from Open AI. 
Secret Key can be generated here: 
``` 
https://platform.openai.com/account/api-keys
 ```

After Generating the key, pass it as an environment varibale like,

``` 
export OPENAI_TOKEN="Your Secret Key"
```

or else you can pass the token with -t flag.

Access it by entering the command 

```
chatgptdash
```
 
 or in case you don't want token in Environment Variable

```
chatgptdash -t "You Secret Key"
```

Output

```
chatgptdash                                                                                                                                 ✔  
Ask to chat gpt: Where is the G20 Hosted?


The G20 is hosted in different countries each year. The host country rotates annually among the G20 members, which include Argentina, Australia, Brazil, Canada, China, France, Germany, India, Indonesia, Italy, Japan, Mexico, Russia, Saudi Arabia, South Africa, South Korea, Turkey, the United Kingdom, the United States, and the European Union. The most recent G20 Summit was hosted by Saudi Arabia in 2020. The upcoming G20 Summit is scheduled to be held in Italy in 2021.

```

## Contributing

Want to make it better or add a feature? Open a pull request and I will review, test and deploy.


