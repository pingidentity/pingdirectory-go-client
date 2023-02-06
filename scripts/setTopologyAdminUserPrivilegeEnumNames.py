#!/usr/bin/env python3
# coding: UTF-8

# The Enum generated for the user-privilege prop for the topology-admin config object is not formatted correctly,
# due to some of the possible values starting with a dash. For example, two possible values are
# "bypass-acl" and "-bypass-acl".
#
# When the generator writes the names for the enum values, it strips any dashes. This leads to duplicate
# enum names like this:
# 	ENUMTOPOLOGYADMINUSERPRIVILEGEPROP_BYPASS_ACL                                 EnumtopologyAdminUserPrivilegeProp = "bypass-acl"
#   ENUMTOPOLOGYADMINUSERPRIVILEGEPROP_BYPASS_ACL                                 EnumtopologyAdminUserPrivilegeProp = "-bypass-acl"
# Causing compile errors.
# 
# To fix this, we add a "REVOKE" before the name of the prop for values starting with a dash. For example:
# 	ENUMTOPOLOGYADMINUSERPRIVILEGEPROP_BYPASS_ACL                                 EnumtopologyAdminUserPrivilegeProp = "bypass-acl"
#   ENUMTOPOLOGYADMINUSERPRIVILEGEPROP_REVOKE_BYPASS_ACL                          EnumtopologyAdminUserPrivilegeProp = "-bypass-acl"

import fileinput

for line in fileinput.input("model_enumtopology_admin_user_privilege_prop.go", inplace=True):
    if "EnumtopologyAdminUserPrivilegeProp = \"-" in line:
        print(line.replace("ENUMTOPOLOGYADMINUSERPRIVILEGEPROP_", "ENUMTOPOLOGYADMINUSERPRIVILEGEPROP_REVOKE_"), end='')
    else:
        print(line, end='')
