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
chatgptdash                                                                                                        
Ask to chat gpt: What is G20 hosted?  


The G20 (Group of Twenty) is a summit of the leaders of the world's largest economies, representing 85% of global GDP (Gross Domestic Product). The host country of the G20 summit rotates annually among its members, who include Argentina, Australia, Brazil, Canada, China, France, Germany, India, Indonesia, Italy, Japan, Mexico, Russia, Saudi Arabia, South Africa, South Korea, Turkey, the United Kingdom, and the United States. The G20 summit tackles a wide range of issues, including global economic growth, trade, financial regulation, climate change, and humanitarian concerns. The host country sets the agenda of the summit, and the meetings are typically attended by high-level officials from the participating countries, including heads of state or government, ministers, and central bank governors.
```

## Contributing

Want to make it better or add a feature? Open a pull request and I will review, test and deploy.


