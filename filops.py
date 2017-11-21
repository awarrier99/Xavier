import crypt

basic = [int, float, bool, str]
basic_str = ["int", "float", "bool", "str"]

def clear(file):
    open(file, "w").close()

def save(file, varname, varval):
    with open(file, "a") as of:
        if type(varval) in basic:
            vtype = basic_str[basic.index(type(varval))]
            varline = vtype + ": " + varname + " = " + str(varval) + "\n"
            enc_varline = crypt.encrypt(varline)
            of.write(enc_varline)

def save_all(file, dict_vars):
    with open(file, "a") as of:
        for ident in dict_vars:
            x = dict_vars[ident]
            save(file, ident, x)

def load(file, varname, line = None):
    with open(file, "r") as inf:
        if line is None:
            l = inf.readlines()
            for enc_line in l:
                line = crypt.decrypt(enc_line)
                if line.find(varname) != -1:
                    break
                
        pos = line.find(":")
        vtype_str = line[:pos]
        if vtype_str in basic_str:
            vtype = basic[basic_str.index(vtype_str)]
            pos = line.find("=") + 2
            varval = vtype(line[pos:].strip())
    return {varname : varval}

def load_all(file):
    dict_vars = {}
    with open(file, "r") as inf:
        l = inf.readlines()
        for enc_line in l:
            line = crypt.decrypt(enc_line)
            varname = line[line.find(":") + 2:line.find("=") - 1]
            x = load(file, varname, line)
            dict_vars.update(x)
    return dict_vars
                

##x = {"var1" : 5, "var2" : "abc", "var3" : True, "var4" : 3.14}
##clear("test")
##save_all("test", x)
##save("test", "var5", "def")
##with open("test", "r") as inf:
##    print(inf.read())
##
##print(load("test", "var2"))
##print(load_all("test"))
