Name: flip
Version: %{_version}
Release: %{_release}%{?dist}
Summary: Floating IP management for Docker Swarm
Group: Tools/Docker

License: GPLv3+

URL: https://github.com/kassisol/flip
Vendor: Kassisol
Packager: Kassisol <support@kassisol.com>

BuildArch: x86_64
BuildRoot: %{_tmppath}/%{name}-buildroot

Source: flip.tar.gz

%description
Flip is an application to manage a Floating IP address on a Docker Swarm cluster.

%prep
%setup -n %{name}

%install
# install binary
install -d $RPM_BUILD_ROOT/%{_sbindir}
install -p -m 755 flip $RPM_BUILD_ROOT/%{_sbindir}/

# add init scripts
install -d $RPM_BUILD_ROOT/%{_unitdir}
install -p -m 644 flip.service $RPM_BUILD_ROOT/%{_unitdir}/flip.service

%files
#%doc README.md LICENSE
%{_sbindir}/flip
/%{_unitdir}/flip.service

%post
%systemd_post flip.service

%preun
%systemd_preun flip.service

%clean
rm -rf $RPM_BUILD_ROOT
