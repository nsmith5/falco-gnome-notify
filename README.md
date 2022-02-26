# Falco Gnome Notify

`falco-gnome-notify` sends [Falco](https://falco.org) alerts to Gnome desktop
notifications. To use it modify your Falco configuration as follows

```diff
# falco.yaml

- json_output: false
+ json_output: true

- json_include_output_property: false
+ json_include_output_property: true

+ program_output:
+   enabled: true
+   keep_alive: true
+   program: "sudo -u <your-username> falco-gnome-notify"
```

## TODO / Known Bugs

- [ ] Don't use `sudo -u` to switch to user
- [ ] Log out and back in breaks notifications
- [ ] Program crash seems to crash falco service as well 
