"""
    Create a CSV string from a JSON Object
    Limitations:
        Can't process JSON Objects inside JSON Arrays
"""
import pandas as pd


def __get_len(odict):
        """
        assumes that all values of a plained dict argument are a list type

        Returns: -> list
            list -> An array which the length of each self.odict value
        """
        each_row_length = {k: len(v) for k, v in odict.items()}
        return each_row_length


def __fill_blank_lines(odict, content=None):
        """
        Receive a plained dict (returned by __normalize())
        and fill differences between child length lines
        """
        _len = __get_len(odict)
        keymax = max(_len, key=_len.get)
        valuemax = _len.pop(keymax)
        for key, _len in _len.items():
            if _len < valuemax:
                for i in range(_len, valuemax):
                    odict[key].append(content)
        return odict


def __to_csv(odict, sep=';'):
        return pd.DataFrame(odict).to_csv(sep=sep)


def __rasterize(p_dict):
    """
    ( Necessary to call __normalize() with the desired dict before )
    Receives a multilevel dict and turns all keys to be 1=1 relationship
    """
    p = {}
    for k, v in p_dict.items():
        if type(v) == dict:
            p = {**p, **__rasterize(v)}
        elif type(v) == list:
            p.__setitem__(k, v)
    return p


def __normalize(odict, father_key=''):
    """
    Transform all child in "father.child" notation
    """
    aux = {}
    if type(odict) is dict:
        for k, v in odict.items():
            key_name = f'{father_key}.{k}' if not father_key == '' else k
            child = {key_name: __normalize(v, key_name)}
            aux = {**aux, **child}
    else:
        return [odict] if not type(odict) is list else odict

    return aux


def to_csv(odict):
    aux = __normalize(dict(odict))
    aux = __rasterize(aux)
    aux = __fill_blank_lines(aux)
    return __to_csv(aux)
