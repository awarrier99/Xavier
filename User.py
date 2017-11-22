class User:
    def __init__(self, first_name, last_name, DOB, nick_name = None, PIN = None):
        self.first_name = first_name
        self.last_name = last_name
        self.DOB = DOB
        self.nick_name = nick_name
        self.PIN = PIN

    def __repr__(self):
        L = [self.first_name, self.last_name, self.DOB, self.nick_name, self.PIN]
        s = "User.User("
        
        for val in L:
            if not val is None:
                if type(val) == str:
                    s += "\""
                    
                s += str(val)
                
                if type(val) == str:
                    s += "\""
                    
                s += ", "

        s = s[:-2]
        s += ")"
        
        return s

    def __str__(self):
        s = "User "
        s += self.first_name + " " + self.last_name
        if not self.nick_name is None:
            s += ", nicknamed " + self.nick_name
        s += ", born on " + self.DOB
        if not self.PIN is None:
            s += ", with PIN " + str(self.PIN)

        return s
