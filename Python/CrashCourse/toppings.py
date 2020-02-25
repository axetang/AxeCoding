requested_toppings = ['mushroom', 'green peppers', 'extra cheese']
if 'mushroom' in requested_toppings:
    print("Adding mushroom")
elif 'pepperoni' in requested_toppings:
    print("adding pepperoni")
elif 'extra cheese' in requested_toppings:
    print("adding extra cheese")

for requested_topping in requested_toppings:
    if requested_topping == 'green peppers':
        print("sorry, we are out of green peppers right now.")
    else:
        print("Adding " + requested_topping+".")

print("\nFinished making your pizza.")

rt = []
if rt:
    for r in rt:
        print("Adding "+r)
else:
    print("R you sure you want a plain pizza?")
