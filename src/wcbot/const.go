package main

// Define constants used by the application
const (
	version = "v0.1.0-alpha.1"

	// Bot command messages:
	// Help message
	msgHelp = "*Wing Commander* here. I will help you to manage and control your SkyMiner and its Nodes.\n\n" +
		"*Usage:*\n" +
		"- /help - show this message\n" +
		"- /about - show information and credits about my creator and any contributors\n" +
		//"- /licences - show licences attribution for original and derived works used by this project\n" +
		"- /status - ask me how I'm going.. and if I'm still running\n" +
		"- /heartbeat - ask me to send you a notification every 2hrs to let you know I’m still running\n" +
		"- /start - start me activly monitoring your SkyMiner. Once started, I will send notifications to you for events that occur\n" +
		"- /stop - stop me monitoring your SkyMiner. Once stopped, I won't send any more notifications\n" +
		"\n" +
		"\n" +
		"Note: I am bound to _this_ chat."

	msgHelpShort = "*Usage:*\n" +
		"- /help - show this message\n" +
		"- /about - show information and credits about my creator and any contributors\n" +
		//"- /licences - show licences attribution for original and derived works used by this project\n" +
		"- /status - ask me how I'm going.. and if I'm still running\n" +
		"- /heartbeat - ask me to send you a notification every 2hrs to let you know I’m still running\n" +
		"- /start - start me activly monitoring your SkyMiner. Once started, I will send notifications to you for events that occur\n" +
		"- /stop - stop me monitoring your SkyMiner. Once stopped, I won't send any more notifications\n"

	// About cmd message
	msgAbout = "*Wing Commander (" + version + ")*\n" +
		"A Telegram Bot written in *Go* designed to help the *Skyfleet* community monitor and manage their *SkyMiners*.\n" +
		"\n" +
		"*Created by:* @BigOokie 2018\n" +
		"*GitHub:* https://github.com/BigOokie/skywire-wing-commander\n" +
		"*Twitter:* https://twitter.com/BigOokie\n" +
		"\n" +
		"Issues and feature requests must be logged via [GitHub](https://github.com/BigOokie/Skywire-Wing-Commander/issues)\n" +
		"\n" +
		"*SkyCoin*: https://www.skycoin.net/\n" +
		"\n" +
		"*Donations most welcome* 👍\n" +
		"*Skycoin:* ES5LccJDhBCK275APmW9tmQNEgiYwTFKQF\n" +
		"*BitCoin:* 37rPeTNjosfydkB4nNNN1XKNrrxxfbLcMA\n"

	// Status cmd message
	msgStatus = "*Wing Commander* Ready and reporting for duty 👍"

	// Start cmd messages
	msgMonitorAlreadyStarted = "*Wing Commander* Sky Manager Monitoring has already been started."
	msgMonitorStart          = "*Wing Commander* Sky Manager Monitoring starting..."

	// Stop cmd message
	msgMonitorStop       = "*Wing Commander* Sky Manager Monitoring stopping..."
	msgMonitorStopped    = "*Wing Commander* Sky Manager Monitoring stopped..."
	msgMonitorNotRunning = "*Wing Commander* Sky Manager Monitoring is not running..."

	// Heartbeat cmd messages
	msgHeartBeatStarted = "*Wing Commander* Heatbeat notifications started. Interval %v"

	/*
		// Licences cmd messages
		msgHandleLicences = "*Wing Commander* Licence Attribution\n" +
			"\n" +
			"**Images and Icons**\n" +
			"- 'Skycoin-Cloud-White@1x.png' by the [SkyCoin Project](https://skycoin.net)\n" +
			"- 'Military Rank' by Ilsur Aptukov from the [Noun Project](https://thenounproject.com/term/military-rank/19966)\n"
	*/

	// Default cmd message (unhandled)
	msgDefault = "Sorry. I don't know that command."

	// OS Interupt Signals
	msgOSInteruptSig = "*Wing Commander* OS Interupt Signal Received. Exiting."
	msgOSKillSig     = "*Wing Commander* OS Kill Signal Received. Exiting."
)
